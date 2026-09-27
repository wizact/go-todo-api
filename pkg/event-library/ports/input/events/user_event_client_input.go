package event

import (
	"context"

	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

type Unsubscribe func() error

type UserEventClientInput interface {
	SubscribeToNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent) (Unsubscribe, error)
}
