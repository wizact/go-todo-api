package event

import (
	"context"

	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserEventClientOutput interface {
	PublishNewUserRegisteredEvent(ctx context.Context, user ua.User) error
}
