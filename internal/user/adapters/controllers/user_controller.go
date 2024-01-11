package controller

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	httpmodel "github.com/wizact/go-todo-api/internal/user/adapters/controllers/models"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	usecase "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type UserController struct {
	userAccountUseCase usecase.UserAccountUseCase
}

func NewUserController(uasuc usecase.UserAccountUseCase) UserController {
	return UserController{
		userAccountUseCase: uasuc,
	}
}

func (u *UserController) RegisterNewUser(ctx context.Context, user httpmodel.User) (httpmodel.User, *hsm.AppError) {
	var ua aggregate.User
	var appErr *hsm.AppError

	// map model to aggregate
	ua, err := user.ToDomainModel()
	if err != nil {
		return user, &hsm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	ua, appErr = u.userAccountUseCase.RegisterNewUser(ctx, ua)

	if appErr != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: appErr, Message: appErr.Error(), Code: http.StatusBadRequest}
	}

	// map aggregate to model
	err = user.ToApiModel(ua)
	if err != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}

func (u *UserController) GetUserById(ctx context.Context, uid uuid.UUID) (httpmodel.User, *hsm.AppError) {
	var ua aggregate.User
	var user httpmodel.User
	var appErr *hsm.AppError

	ua, appErr = u.userAccountUseCase.GetUserById(ctx, uid)

	if appErr != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: appErr, Message: appErr.Error(), Code: http.StatusBadRequest}
	}

	// map aggregate to model
	err := user.ToApiModel(ua)
	if err != nil {
		// return proper error
		return user, &hsm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}
