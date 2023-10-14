package service

import (
	"context"

	repository "github.com/wizact/go-todo-api/internal/user/domain/repository"
)

type EmailVerificationService struct {
	// repositories and other services
	emailGatewayRepository repository.EmailGatewayRepository
	Flag                   int
}

func NewEmailVerificationService(egr repository.EmailGatewayRepository) *EmailVerificationService {
	evs := &EmailVerificationService{emailGatewayRepository: egr}

	return evs
}

func (ev *EmailVerificationService) SendEmailVerificationEmail(email string) error {
	return ev.emailGatewayRepository.SendEmail(context.Background(), email, "Please verify your email", "Please click on this link to verify your email: link")
}
