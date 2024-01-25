package events

import (
	"context"

	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserEventClient interface {
	UserCreated(ctx context.Context, user ua.User) error
}