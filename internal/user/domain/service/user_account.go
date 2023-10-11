package service

import (
	"context"
	"errors"

	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	repository "github.com/wizact/go-todo-api/internal/user/domain/repository"
)

type UserAccountService struct {
	// repositories and other services
	userRepository           repository.UserRepository
	emailVerificationService EmailVerificationService
}

func NewUserAccountService(cfgs ...UserAccountServiceConfigurations) (*UserAccountService, error) {
	ua := &UserAccountService{}
	for _, cfgs := range cfgs {
		err := cfgs(ua)
		if err != nil {
			return nil, err
		}
	}

	return ua, nil
}

func (ua *UserAccountService) RegisterNewUser(user aggregate.User) (*aggregate.User, error) {
	// Verify the account
	if !user.IsValid() {
		return nil, errors.New("user info is not valid")
	}

	// Check if the user does not exist
	u, e := ua.userRepository.FindByEmail(context.Background(), user.Email())
	if e != nil {
		return nil, e
	}

	if u != nil {
		return nil, errors.New("email already registered")
	}

	// Create user
	u, e = ua.userRepository.Create(context.Background(), user)
	if e != nil {
		return nil, e
	}

	// Trigger new email message (dependency in the service layer to the infra)
	ua.emailVerificationService.SendEmailVerificationEmail(user.Email())
	if e != nil {
		return nil, e
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
