package event

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	pubsub_infra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserEventClientMock struct {
}

func (uecm UserEventClientMock) Connection(nc *pubsub_infra.NatsConnection) {
	log.Println("UserEventClientMock.Connection")
}

func (uecm UserEventClientMock) GetConnection() *pubsub_infra.NatsConnection {
	log.Println("UserEventClientMock.GetConnection")
	return nil
}

// PublishNewUserRegisteredEvent publishes the user aggregate after successful user creation
func (uecm UserEventClientMock) PublishNewUserRegisteredEvent(ctx context.Context, user ua.User) error {
	log.Println("UserEventClientMock.UserCreated")
	return nil
}

func (uecm UserEventClientMock) SubscribeToNewUserRegisteredEvent(ctx context.Context, ch chan *nats.Msg) (pubsub_infra.ChannelUnsubscribeCallBack, error) {
	return nil, nil
}
