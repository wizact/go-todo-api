package controller

import (
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/service"
)

type UserController struct {
	userAccountService usecase.UserAccountUseCase
}

func NewUserController(uauc usecase.UserAccountUseCase) UserController {
	return UserController{
		userAccountService: uauc,
	}
}

func (u *UserController) RegisterNewUser( /* model User */ ) {
	// map model to aggregate
	var ua aggregate.User
	var err error
	_, err = u.userAccountService.RegisterNewUser(ua)

	if err != nil {
		// return proper error
		return
	}

	// map aggregate to model
	// return the result
}
