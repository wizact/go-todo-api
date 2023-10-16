package service

import (
	"context"
	"errors"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	repository "github.com/wizact/go-todo-api/internal/user/interfaces/repository"
)

type UserAccountUseCase interface {
	RegisterNewUser(user aggregate.User) (aggregate.User, error)
	AuthenticateUser(email, password string) (bool, error)
	ValidateNewUserEmail(u aggregate.User) error
}

type UserAccountService struct {
	// repositories and other services
	userRepository           repository.UserRepository
	emailVerificationService EmailVerificationUseCase
}

func NewUserAccountService(ur repository.UserRepository, evs EmailVerificationUseCase) UserAccountUseCase {
	ua := &UserAccountService{
		userRepository:           ur,
		emailVerificationService: evs,
	}

	return ua
}

func (ua *UserAccountService) RegisterNewUser(user aggregate.User) (aggregate.User, error) {
	// Verify the account
	if !user.IsValid() {
		return user, errors.New("user info is not valid")
	}

	// Check if the user does not exist
	u, e := ua.userRepository.FindByEmail(context.Background(), user.Email())
	if e != nil {
		return user, e
	}

	if u.Email() != user.Email() {
		return user, errors.New("email already registered")
	}

	// Create user
	u, e = ua.userRepository.Create(context.Background(), user)
	if e != nil {
		return user, e
	}

	// Trigger new email message (dependency in the service layer to the infra)
	ua.emailVerificationService.SendEmailVerificationEmail(user.Email())
	if e != nil {
		return user, e
	}

	return u, nil
}

func (ua *UserAccountService) ValidateNewUserEmail(u aggregate.User) error {
	// Verify the token
	// Set the flag to enable the user
	return nil
}

func (ua *UserAccountService) AuthenticateUser(email, password string) (bool, error) {
	// Verify the token
	// Set the flag to enable the user
	return true, nil
}
