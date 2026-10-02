package controller

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	httpmodel "github.com/wizact/go-todo-api/internal/user/adapters/controllers/models"
	"github.com/wizact/go-todo-api/internal/user/domain"
	userAppSvc "github.com/wizact/go-todo-api/internal/user/ports/applications"
	usecase "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type UserController struct {
	userAccountUseCase usecase.UserAccountUseCase
	registrationAppSvc userAppSvc.Registration
}

func NewUserController(uasuc usecase.UserAccountUseCase, rappsvc userAppSvc.Registration) UserController {
	return UserController{
		userAccountUseCase: uasuc,
		registrationAppSvc: rappsvc,
	}
}

func (u *UserController) VerifyUserRegistration(ctx context.Context, uid uuid.UUID, token string) *hsm.AppError {
	// TODO: AuthZ check (own user or admin)
	err := u.registrationAppSvc.VerifyUserRegistration(ctx, uid, token)

	if err != nil {
		log.Println(err)
		return hsm.NewAppError(err, "Failed activating the user", http.StatusBadRequest)
	}

	return nil
}

func (u *UserController) RegisterNewUser(ctx context.Context, user httpmodel.User) (httpmodel.User, *hsm.AppError) {
	// map model to aggregate
	ua, appErr := user.ToDomainModel()
	if appErr != nil {
		return user, &hsm.AppError{ErrorObject: appErr, SanitisedMessage: appErr.Error(), Code: http.StatusBadRequest}
	}

	ua, err := u.userAccountUseCase.RegisterNewUser(ctx, ua)

	if err != nil {
		log.Println(err)
		return user, userAccountRegistrationAppError(err)
	}

	// map aggregate to model
	appErr = user.ToApiModel(ua)
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

func (u *UserController) GetUserById(ctx context.Context, uid uuid.UUID) (httpmodel.User, *hsm.AppError) {
	// TODO: AuthZ check (own user or admin)
	var user httpmodel.User

	ua, err := u.userAccountUseCase.GetUserById(ctx, uid)

	if err != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: err, SanitisedMessage: err.Error(), Code: http.StatusBadRequest}
	}

	// map aggregate to model
	appErr := user.ToApiModel(ua)
	if appErr != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: appErr, SanitisedMessage: appErr.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}
