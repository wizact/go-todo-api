package service

import (
	"context"

	"github.com/google/uuid"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -source=registration.go -destination=../mocks/registration.go -package=mocks

type Registration interface {
	GetRegistrationVerificationEmailData(uid uuid.UUID) (map[string]string, error)
	VerifyUserRegistration(ctx context.Context, uid uuid.UUID, hash string) *hsm.AppError
	Done()
}
