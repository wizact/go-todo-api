package service

// Email verification service factory can tie repo interface with infra concrete implementation.
import (
	repository "github.com/wizact/go-todo-api/internal/user/domain/repository"
)

type EmailVerificationServiceConfigurations func(ua *EmailVerificationService) error

func WithEmailGatewayRepository(egr repository.EmailGatewayRepository) EmailVerificationServiceConfigurations {
	return func(ev *EmailVerificationService) error {
		ev.emailGatewayRepository = egr
		return nil
	}
}
