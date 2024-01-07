package service

import (
	"context"
	"errors"
	"net/http"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"github.com/wizact/go-todo-api/internal/user/domain/model"
	repository "github.com/wizact/go-todo-api/internal/user/interfaces/repository"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

var (
	ErrInvalidUser          = hsm.NewAppError(errors.New("user info is not valid"), "user info is not valid", http.StatusBadRequest)
	ErrFailedToRegisterUser = hsm.NewAppError(errors.New("user info is not valid"), "user info is not valid", http.StatusBadRequest)
	ErrEmailAlreadyExists   = hsm.NewAppError(errors.New("email already registered"), "email already registered", http.StatusBadRequest)
)

type UserAccountUseCase interface {
	RegisterNewUser(user aggregate.User) (aggregate.User, *hsm.AppError)
	AuthenticateUser(email, password string) (bool, *hsm.AppError)
	ValidateNewUserEmail(u aggregate.User) *hsm.AppError
}

type UserAccountService struct {
	// repositories and other services
	userRepository repository.UserRepository
}

func NewUserAccountService(ur repository.UserRepository) UserAccountUseCase {
	ua := &UserAccountService{
		userRepository: ur,
	}

	return ua
}

func (ua *UserAccountService) RegisterNewUser(user aggregate.User) (aggregate.User, *hsm.AppError) {
	// Set the new user role to Limited
	r := user.Role()
	r.Name = model.Limited
	user.SetRole(r)

	// Verify the account
	if !user.IsValid() {
		return user, ErrInvalidUser
	}

	// Check if the user does not exist
	u, e := ua.userRepository.FindByEmail(context.Background(), user.Email())
	if e != nil {
		return user, ErrFailedToRegisterUser
	}

	if u.Email() == user.Email() {
		return user, ErrEmailAlreadyExists
	}

	// Create user
	u, e = ua.userRepository.Create(context.Background(), user)
	if e != nil {
		return user, ErrFailedToRegisterUser
	}

	return u, nil
}

func (ua *UserAccountService) ValidateNewUserEmail(u aggregate.User) *hsm.AppError {
	// Verify the token
	// Set the flag to enable the user
	return nil
}

func (ua *UserAccountService) AuthenticateUser(email, password string) (bool, *hsm.AppError) {
	// Verify the token
	// Set the flag to enable the user
	return true, nil
}
