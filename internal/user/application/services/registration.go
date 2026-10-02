package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	applicationregistration "github.com/wizact/go-todo-api/internal/user/application/registration"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	repository_port "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

const registrationVerificationLifetime = 24 * time.Hour

// Registration application service responsible for managing the lifecycle of a user registration
type Registration struct {
	userAccountUseCase     usecase_port.UserAccountUseCase
	registrationRepository repository_port.RegistrationRepository
	now                    func() time.Time
	done                   chan bool
}

// NewRegistration returns a new registration application service.
func NewRegistration(
	uc usecase_port.UserAccountUseCase,
	repository repository_port.RegistrationRepository,
) *Registration {
	return &Registration{
		userAccountUseCase:     uc,
		registrationRepository: repository,
		now:                    time.Now,
		done:                   make(chan bool),
	}
}

func (r *Registration) Done() {
	r.done <- true
}

// GetRegistrationVerificationEmailData returns the data required to send a registration verification email
func (r *Registration) GetRegistrationVerificationEmailData(uid uuid.UUID) (map[string]string, error) {
	em := make(map[string]string)
	u, err := r.userAccountUseCase.GetUserById(context.Background(), uid)
	if err != nil {
		return em, fmt.Errorf("get user for registration verification email: %w", err)
	}

	token, digest, err := issueRegistrationVerification()
	if err != nil {
		return em, fmt.Errorf("issue registration verification: %w", err)
	}
	verification := applicationregistration.Verification{
		UserID:       uid,
		SecretDigest: digest,
		ExpiresAt:    r.now().UTC().Add(registrationVerificationLifetime),
	}
	if err := r.registrationRepository.SaveRegistrationVerification(context.Background(), verification); err != nil {
		return em, fmt.Errorf("save registration verification: %w", err)
	}

	ue := u.User()
	em["email"] = u.Email()
	em["nick_name"] = ue.ConcatenatedName()
	em["token"] = token
	em["base_url"] = "http://localhost:8080" //TODO: get base url from env
	em["verify_email_link"] = fmt.Sprintf("%s/users/verify-registration?uid=%s&token=%s", em["base_url"], uid.String(), token)

	return em, nil
}

func (r *Registration) VerifyUserRegistration(ctx context.Context, uid uuid.UUID, token string) error {
	verification, err := r.registrationRepository.FindRegistrationVerification(ctx, uid)
	if err != nil {
		return fmt.Errorf("find registration verification: %w", err)
	}
	if !r.now().UTC().Before(verification.ExpiresAt) {
		return applicationregistration.ErrVerificationExpired
	}
	if !matchesRegistrationVerification(verification.SecretDigest, token) {
		return applicationregistration.ErrVerificationSecretMismatch
	}

	u, err := r.userAccountUseCase.GetUserById(ctx, uid)
	if err != nil {
		return fmt.Errorf("get user for registration verification: %w", err)
	}

	u.VerifyRegistration()

	if _, err := r.registrationRepository.CompleteRegistration(ctx, u, verification.SecretDigest); err != nil {
		return fmt.Errorf("complete user registration: %w", err)
	}

	return nil
}
