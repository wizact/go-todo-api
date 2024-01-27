package event

import (
	"context"
	"log"

	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserEventClientMock struct {
}

func (uecm UserEventClientMock) Connection(nc *pubsubinfra.NatsConnection) {
	log.Println("UserEventClientMock.Connection")
}

func (uecm UserEventClientMock) GetConnection() *pubsubinfra.NatsConnection {
	log.Println("UserEventClientMock.GetConnection")
	return nil
}

// NewUserRegistered publishes the user aggregate after successful user creation
func (uecm UserEventClientMock) NewUserRegistered(ctx context.Context, user ua.User) error {
	log.Println("UserEventClientMock.UserCreated")
	return nil
}
