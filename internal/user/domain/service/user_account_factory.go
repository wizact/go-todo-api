package service

// User account factory can tie repo interface with infra concrete implementation, hence is the only place it should cross the hexagonal layers
import (
	repository "github.com/wizact/go-todo-api/internal/user/domain/repository"
)

type UserAccountServiceConfigurations func(ua *UserAccountService) error

func WithUserSqlLiteRepository(ur repository.UserRepository) UserAccountServiceConfigurations {
	return func(ua *UserAccountService) error {
		ua.userRepository = ur
		return nil
	}
}

// WithCustomUserRepository enables to provide a custom repo (i.e. memory repo). Can be used for tests and still decouple service from infra layer
func WithCustomUserRepository(customUserRepo repository.UserRepository) UserAccountServiceConfigurations {
	return WithUserSqlLiteRepository(customUserRepo)
}

func WithEmailVerificationService(evs EmailVerificationService) UserAccountServiceConfigurations {
	return func(ua *UserAccountService) error {
		ua.emailVerificationService = evs
		return nil
	}
}
