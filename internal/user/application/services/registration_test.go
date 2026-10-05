package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	repositoryadapter "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	applicationregistration "github.com/wizact/go-todo-api/internal/user/application/registration"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	"github.com/wizact/go-todo-api/internal/user/ports/mocks"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

type registrationIssuanceState struct {
	TokenPresent          bool
	LegacyHashAbsent      bool
	Link                  string
	DigestMatches         bool
	ExpiresAt             time.Time
	RepositoryReceivedCtx bool
}

type registrationVerificationState struct {
	ErrorMatches       bool
	Active             bool
	EmailVerified      bool
	CredentialRetained bool
	ReplayWasRejected  bool
}

func TestRegistration_GetRegistrationVerificationEmailData_IssuesStoredCredential(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	user := registrationTestUser(userID)
	requestContext := context.WithValue(context.Background(), registrationContextKey{}, "issuance")
	now := time.Date(2026, time.October, 3, 12, 0, 0, 0, time.UTC)
	memoryRepository := repositoryadapter.NewUserMemoryRepository(nil)
	repository := &registrationRepositoryContextSpy{RegistrationRepository: memoryRepository}
	userAccount.EXPECT().
		GetUserById(requestContext, userID).
		Return(user, nil)
	registration := NewRegistration(userAccount, repository)
	registration.now = func() time.Time { return now }

	emailData, err := registration.GetRegistrationVerificationEmailData(requestContext, userID)
	if err != nil {
		t.Fatalf("GetRegistrationVerificationEmailData() error = %v", err)
	}
	verification, err := memoryRepository.FindRegistrationVerification(context.Background(), userID)
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}
	token := emailData["token"]

	got := registrationIssuanceState{
		TokenPresent:          token != "",
		LegacyHashAbsent:      emailData["hash"] == "",
		Link:                  emailData["verify_email_link"],
		DigestMatches:         matchesRegistrationVerification(verification.SecretDigest, token),
		ExpiresAt:             verification.ExpiresAt,
		RepositoryReceivedCtx: repository.saveContext == requestContext,
	}
	want := registrationIssuanceState{
		TokenPresent:          true,
		LegacyHashAbsent:      true,
		Link:                  "http://localhost:8080/users/verify-registration?uid=" + userID.String() + "&token=" + token,
		DigestMatches:         true,
		ExpiresAt:             now.Add(24 * time.Hour),
		RepositoryReceivedCtx: true,
	}
	if got != want {
		t.Fatalf("registration issuance state = %#v, want %#v", got, want)
	}
}

type registrationContextKey struct{}

type registrationRepositoryContextSpy struct {
	repositoryport.RegistrationRepository
	saveContext context.Context
}

func (repository *registrationRepositoryContextSpy) SaveRegistrationVerification(
	ctx context.Context,
	verification applicationregistration.Verification,
) error {
	repository.saveContext = ctx
	return repository.RegistrationRepository.SaveRegistrationVerification(ctx, verification)
}

func TestRegistration_VerifyUserRegistration_ValidTokenCompletesRegistration(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	user := registrationTestUser(userID)
	now := time.Date(2026, time.October, 3, 12, 0, 0, 0, time.UTC)
	token := "valid-registration-token"
	repository := repositoryadapter.NewUserMemoryRepository([]aggregate.User{user})
	storeRegistrationVerification(t, repository, userID, token, now.Add(time.Hour))
	userAccount.EXPECT().
		GetUserById(gomock.Any(), userID).
		Return(user, nil)
	registration := NewRegistration(userAccount, repository)
	registration.now = func() time.Time { return now }

	err := registration.VerifyUserRegistration(context.Background(), userID, token)
	replayError := registration.VerifyUserRegistration(context.Background(), userID, token)
	persistedUser, findUserError := repository.FindById(context.Background(), userID)
	if findUserError != nil {
		t.Fatalf("FindById() error = %v", findUserError)
	}
	_, credentialError := repository.FindRegistrationVerification(context.Background(), userID)

	got := registrationVerificationState{
		ErrorMatches:       err == nil,
		Active:             persistedUser.IsActive(),
		EmailVerified:      persistedUser.HasVerifiedEmail(),
		CredentialRetained: !errors.Is(credentialError, repositoryport.ErrRegistrationVerificationNotFound),
		ReplayWasRejected:  errors.Is(replayError, repositoryport.ErrRegistrationVerificationNotFound),
	}
	want := registrationVerificationState{
		ErrorMatches:      true,
		Active:            true,
		EmailVerified:     true,
		ReplayWasRejected: true,
	}
	if got != want {
		t.Fatalf("registration verification state = %#v, want %#v", got, want)
	}
}

func TestRegistration_VerifyUserRegistration_WrongTokenPreservesRegistration(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	user := registrationTestUser(userID)
	now := time.Date(2026, time.October, 3, 12, 0, 0, 0, time.UTC)
	repository := repositoryadapter.NewUserMemoryRepository([]aggregate.User{user})
	verification := storeRegistrationVerification(t, repository, userID, "valid-token", now.Add(time.Hour))
	registration := NewRegistration(userAccount, repository)
	registration.now = func() time.Time { return now }

	err := registration.VerifyUserRegistration(context.Background(), userID, "wrong-token")
	persistedUser, findUserError := repository.FindById(context.Background(), userID)
	if findUserError != nil {
		t.Fatalf("FindById() error = %v", findUserError)
	}
	persistedVerification, findVerificationError := repository.FindRegistrationVerification(context.Background(), userID)
	if findVerificationError != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", findVerificationError)
	}

	got := registrationVerificationState{
		ErrorMatches:       errors.Is(err, applicationregistration.ErrVerificationSecretMismatch),
		Active:             persistedUser.IsActive(),
		EmailVerified:      persistedUser.HasVerifiedEmail(),
		CredentialRetained: persistedVerification == verification,
	}
	want := registrationVerificationState{ErrorMatches: true, CredentialRetained: true}
	if got != want {
		t.Fatalf("registration verification state = %#v, want %#v", got, want)
	}
}

func TestRegistration_VerifyUserRegistration_ExpiredTokenPreservesRegistration(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	user := registrationTestUser(userID)
	now := time.Date(2026, time.October, 3, 12, 0, 0, 0, time.UTC)
	repository := repositoryadapter.NewUserMemoryRepository([]aggregate.User{user})
	verification := storeRegistrationVerification(t, repository, userID, "expired-token", now)
	registration := NewRegistration(userAccount, repository)
	registration.now = func() time.Time { return now }

	err := registration.VerifyUserRegistration(context.Background(), userID, "expired-token")
	persistedUser, findUserError := repository.FindById(context.Background(), userID)
	if findUserError != nil {
		t.Fatalf("FindById() error = %v", findUserError)
	}
	persistedVerification, findVerificationError := repository.FindRegistrationVerification(context.Background(), userID)
	if findVerificationError != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", findVerificationError)
	}

	got := registrationVerificationState{
		ErrorMatches:       errors.Is(err, applicationregistration.ErrVerificationExpired),
		Active:             persistedUser.IsActive(),
		EmailVerified:      persistedUser.HasVerifiedEmail(),
		CredentialRetained: persistedVerification == verification,
	}
	want := registrationVerificationState{ErrorMatches: true, CredentialRetained: true}
	if got != want {
		t.Fatalf("registration verification state = %#v, want %#v", got, want)
	}
}

func TestRegistration_VerifyUserRegistration_MissingCredentialReturnsNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	repository := repositoryadapter.NewUserMemoryRepository(nil)
	registration := NewRegistration(userAccount, repository)

	err := registration.VerifyUserRegistration(context.Background(), uuid.New(), "missing-token")

	if !errors.Is(err, repositoryport.ErrRegistrationVerificationNotFound) {
		t.Fatalf("VerifyUserRegistration() error = %v, want %v", err, repositoryport.ErrRegistrationVerificationNotFound)
	}
}

func registrationTestUser(userID uuid.UUID) aggregate.User {
	phone := model.NewPhoneNumber("+64", "23", "123456")
	userModel := model.RehydrateUser(
		userID,
		"Foo",
		"Bar",
		time.Date(1983, time.January, 2, 15, 4, 5, 0, time.UTC),
		"foo.bar@example.com",
		phone,
	)
	location := model.NewLocation(173.3002574488138, -41.26595602617756)

	user := aggregate.NewUser()
	user.SetUser(userModel)
	user.SetLocation(location)
	return user
}

func storeRegistrationVerification(
	t *testing.T,
	repository repositoryport.RegistrationRepository,
	userID uuid.UUID,
	token string,
	expiresAt time.Time,
) applicationregistration.Verification {
	t.Helper()

	digest, err := hashRegistrationVerificationSecret(token)
	if err != nil {
		t.Fatalf("hash registration verification secret: %v", err)
	}
	verification := applicationregistration.Verification{
		UserID:       userID,
		SecretDigest: digest,
		ExpiresAt:    expiresAt,
	}
	if err := repository.SaveRegistrationVerification(context.Background(), verification); err != nil {
		t.Fatalf("SaveRegistrationVerification() error = %v", err)
	}

	return verification
}
