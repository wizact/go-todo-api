package event

import (
	"context"
	"log"

	event_input "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	pubsub "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

type UserEventClientMock struct {
}

func (uecm UserEventClientMock) Connection(nc *pubsub.NatsConnection) {
	log.Println("UserEventClientMock.Connection")
}

func (uecm UserEventClientMock) GetConnection() *pubsub.NatsConnection {
	log.Println("UserEventClientMock.GetConnection")
	return nil
}

// PublishNewUserRegisteredEvent publishes the user aggregate after successful user creation
func (uecm UserEventClientMock) PublishNewUserRegisteredEvent(ctx context.Context, userDE ude.UserDomainEvent) error {
	log.Println("UserEventClientMock.UserCreated")
	return nil
}

func (uecm UserEventClientMock) SubscribeToNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent) (event_input.Unsubscribe, error) {
	return func() error { return nil }, nil
}
