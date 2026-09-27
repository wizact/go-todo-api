package repository

import (
	"context"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

// UserRepository is the sole write boundary for the user aggregate and its projections.
// Implementations with writable projections must persist them atomically with the aggregate.
type UserRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (ua.User, error)
	FindByEmail(ctx context.Context, email string) (ua.User, error)
	Create(ctx context.Context, user ua.User) (ua.User, error)
	Update(ctx context.Context, user ua.User) (ua.User, error)
}
