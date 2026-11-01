package event

import (
	"context"

	ude "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
)

type Unsubscribe func() error

type UserSubscriber interface {
	SubscribeToNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent) (Unsubscribe, error)
}
