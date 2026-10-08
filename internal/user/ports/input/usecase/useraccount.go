package usecase

import (
	"context"

	"github.com/google/uuid"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
)

//go:generate go run go.uber.org/mock/mockgen@v0.6.0 -source=useraccount.go -destination=../../mocks/useraccount.go -package=mocks

type UserAccountUseCase interface {
	RegisterNewUser(ctx context.Context, user aggregate.User) (aggregate.User, error)
	FetchUserByID(ctx context.Context, uid uuid.UUID) (aggregate.User, error)
	UpdateUser(ctx context.Context, user aggregate.User) (aggregate.User, error)
}
