package controller

import (
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/service"
)

type UserController struct {
	userAccountUseCase usecase.UserAccountUseCase
}

func NewUserController(uasuc usecase.UserAccountUseCase) UserController {
	return UserController{
		userAccountUseCase: uasuc,
	}
}

func (u *UserController) RegisterNewUser( /* model User */ ) {
	// map model to aggregate
	var ua aggregate.User
	var err error
	_, err = u.userAccountUseCase.RegisterNewUser(ua)

	if err != nil {
		// return proper error
		return
	}

	// map aggregate to model
	// return the result
}
