package event

import (
	"context"

	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

type UserEventClientOutput interface {
	PublishNewUserRegisteredEvent(ctx context.Context, userDe ude.UserDomainEvent) error
}
