package controller

import (
	"net/http"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/service"
	httpmodel "github.com/wizact/go-todo-api/internal/user/interfaces/model"
	hm "github.com/wizact/go-todo-api/pkg/http-model"
)

type UserController struct {
	userAccountUseCase usecase.UserAccountUseCase
}

func NewUserController(uasuc usecase.UserAccountUseCase) UserController {
	return UserController{
		userAccountUseCase: uasuc,
	}
}

func (u *UserController) RegisterNewUser(user httpmodel.User) (httpmodel.User, *hm.AppError) {
	var ua aggregate.User

	// map model to aggregate
	ua, err := user.ToDomainModel()
	if err != nil {
		return user, &hm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	ua, err = u.userAccountUseCase.RegisterNewUser(ua)

	if err != nil {
		// return proper error
		return user, &hm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	// map aggregate to model
	err = user.ToApiModel(ua)
	if err != nil {
		// return proper error
		return user, &hm.AppError{ErrorObject: err, Message: err.Error(), Code: http.StatusBadRequest}
	}

	return user, nil
}
