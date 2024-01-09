package usecase

import (
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type UserAccountUseCase interface {
	RegisterNewUser(user aggregate.User) (aggregate.User, *hsm.AppError)
	AuthenticateUser(email, password string) (bool, *hsm.AppError)
	ValidateNewUserEmail(u aggregate.User) *hsm.AppError
}
