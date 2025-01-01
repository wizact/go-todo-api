package event

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
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

func (uecm UserEventClientMock) SubscribeToNewUserRegisteredEvent(ctx context.Context, ch chan *nats.Msg) (pubsub.ChannelUnsubscribeCallBack, error) {
	return nil, nil
}
