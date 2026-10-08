package controller

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	httpmodel "github.com/wizact/go-todo-api/internal/user/adapters/controller/model"
	"github.com/wizact/go-todo-api/internal/user/domain"
	usecase "github.com/wizact/go-todo-api/internal/user/ports/input/usecase"
	userAppSvc "github.com/wizact/go-todo-api/internal/user/ports/service"
	hsm "github.com/wizact/go-todo-api/pkg/httpservermodel"
)

type Controller struct {
	userAccountUseCase usecase.UserAccountUseCase
	registrationAppSvc userAppSvc.Registration
}

func New(uasuc usecase.UserAccountUseCase, rappsvc userAppSvc.Registration) Controller {
	return Controller{
		userAccountUseCase: uasuc,
		registrationAppSvc: rappsvc,
	}
}

func (c *Controller) VerifyUserRegistration(ctx context.Context, uid uuid.UUID, token string) *hsm.AppError {
	// TODO: AuthZ check (own user or admin)
	err := c.registrationAppSvc.VerifyUserRegistration(ctx, uid, token)

	if err != nil {
		log.Println(err)
		return hsm.NewAppError(err, "Failed activating the user", http.StatusBadRequest)
	}

	return nil
}

func (c *Controller) RegisterNewUser(ctx context.Context, user httpmodel.User) (httpmodel.User, *hsm.AppError) {
	// map model to aggregate
	ua, appErr := user.ToDomainModel()
	if appErr != nil {
		return user, &hsm.AppError{ErrorObject: appErr, SanitisedMessage: appErr.Error(), Code: http.StatusBadRequest}
	}

	ua, err := c.userAccountUseCase.RegisterNewUser(ctx, ua)

	if err != nil {
		log.Println(err)
		return user, userAccountRegistrationAppError(err)
	}

	// map aggregate to model
	appErr = user.ToAPIModel(ua)
	if appErr != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: appErr, SanitisedMessage: appErr.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}

func userAccountRegistrationAppError(err error) *hsm.AppError {
	switch {
	case errors.Is(err, domain.ErrInvalidUser), errors.Is(err, domain.ErrRegistrationFailed):
		return hsm.NewAppError(err, "user info is not valid", http.StatusBadRequest)
	case errors.Is(err, domain.ErrEmailAlreadyExists):
		return hsm.NewAppError(err, "email already registered", http.StatusBadRequest)
	default:
		return hsm.NewAppError(err, "internal server error", http.StatusInternalServerError)
	}
}

func (c *Controller) FetchUserByID(ctx context.Context, uid uuid.UUID) (httpmodel.User, *hsm.AppError) {
	// TODO: AuthZ check (own user or admin)
	var user httpmodel.User

	ua, err := c.userAccountUseCase.FetchUserByID(ctx, uid)

	if err != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: err, SanitisedMessage: err.Error(), Code: http.StatusBadRequest}
	}

	// map aggregate to model
	appErr := user.ToAPIModel(ua)
	if appErr != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: appErr, SanitisedMessage: appErr.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}
