package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/events"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

// Registration application service responsible for managing the lifecycle of a user registration
type Registration struct {
	userEventClient event_port.UserEventClient
	// app service can reference domain service ( but not the other way arround)
	userAccountUseCase usecase_port.UserAccountUseCase
	done               chan bool
}

// NewRegisteration returns a new instance of Registration application service
func NewRegisteration(uec event_port.UserEventClient, uc usecase_port.UserAccountUseCase) *Registration {
	return &Registration{
		userEventClient:    uec,
		userAccountUseCase: uc,
		done:               make(chan bool),
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
		return em, fmt.Errorf("user registration app service > send email verification: %v", err)
	}

	t := u.Token()
	h, e := t.CreateTokenVerificationHash()
	if e != nil {
		return em, fmt.Errorf("user registration app service > hash function failed: %v", e)
	}

	ue := u.User()
	em["email"] = u.Email()
	em["nick_name"] = ue.ConcatenatedName()
	em["hash"] = string(h)
	em["base_url"] = "http://localhost:8080" //TODO: get base url from env
	em["verify_email_link"] = fmt.Sprintf("%s/users/verify-registration?uid=%s&hash=%s", em["base_url"], uid.String(), h)

	return em, nil
}

func (r *Registration) VerifyUserRegistration(ctx context.Context, uid uuid.UUID, hash string) *hsm.AppError {
	u, err := r.userAccountUseCase.GetUserById(ctx, uid)
	if err != nil {
		return &hsm.AppError{ErrorObject: err.ErrorObject, SanitisedMessage: "Failed activating the user", Code: http.StatusBadRequest}
	}

	t := u.Token()
	if !t.CompareTokenVerificationWithHash([]byte(hash)) {
		return &hsm.AppError{ErrorObject: errors.New("hash is not matched"), SanitisedMessage: "Failed activating the user", Code: http.StatusBadRequest}
	}

	u.SetIsActive(true)
	u.SetHasVerifiedEmail(true)

	_, e := r.userAccountUseCase.UpdateUser(ctx, u)
	if e != nil {
		return &hsm.AppError{ErrorObject: e.ErrorObject, SanitisedMessage: "Failed activating the user", Code: http.StatusBadRequest}
	}

	return nil
}
