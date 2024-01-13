package usecase

import (
	"context"

	"github.com/google/uuid"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type UserAccountUseCase interface {
	RegisterNewUser(ctx context.Context, user aggregate.User) (aggregate.User, *hsm.AppError)
	GetUserById(ctx context.Context, uid uuid.UUID) (aggregate.User, *hsm.AppError)
}
