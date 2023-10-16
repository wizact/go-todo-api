package service

import (
	"context"

	repository "github.com/wizact/go-todo-api/internal/user/interfaces/repository"
)

type EmailVerificationUseCase interface {
	SendEmailVerificationEmail(email string) error
}

type EmailVerificationService struct {
	// repositories and other services
	emailGatewayRepository repository.EmailGatewayRepository
}

func NewEmailVerificationService(egr repository.EmailGatewayRepository) EmailVerificationUseCase {
	evs := EmailVerificationService{emailGatewayRepository: egr}

	return &evs
}

func (ev *EmailVerificationService) SendEmailVerificationEmail(email string) error {
	return ev.emailGatewayRepository.SendEmail(context.Background(), email, "Please verify your email", "Please click on this link to verify your email: link")
}
