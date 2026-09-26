package usecase

import (
	"context"

	"github.com/google/uuid"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -source=user_account.go -destination=../../mocks/user_account.go -package=mocks

type UserAccountUseCase interface {
	RegisterNewUser(ctx context.Context, user aggregate.User) (aggregate.User, *hsm.AppError)
	GetUserById(ctx context.Context, uid uuid.UUID) (aggregate.User, *hsm.AppError)
	UpdateUser(ctx context.Context, user aggregate.User) (aggregate.User, *hsm.AppError)
}
