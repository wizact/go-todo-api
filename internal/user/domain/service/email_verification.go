package service

import (
	"context"

	repository "github.com/wizact/go-todo-api/internal/user/domain/repository"
)

type EmailVerificationService struct {
	// repositories and other services
	emailGatewayRepository repository.EmailGatewayRepository
}

func NewEmailVerificationService(cfgs ...EmailVerificationServiceConfigurations) (*EmailVerificationService, error) {
	evs := &EmailVerificationService{}
	for _, cfgs := range cfgs {
		err := cfgs(evs)
		if err != nil {
			return nil, err
		}
	}

	return evs, nil
}

func (ev *EmailVerificationService) SendEmailVerificationEmail(email string) error {
	return ev.emailGatewayRepository.SendEmail(context.Background(), email, "Please verify your email", "Please click on this link to verify your email: link")
}
