package event

import (
	"context"

	"github.com/wizact/go-todo-api/internal/user/domain"
)

type UserEventPublisher interface {
	PublishNewUserRegisteredEvent(ctx context.Context, event domain.UserRegisteredEvent) error
}
