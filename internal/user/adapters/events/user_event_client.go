package events

import (
	"context"
	"encoding/json"

	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

const UserCreatedEventSubject = "UserCreated"

type UserEventClient struct {
	natConnection *pubsubinfra.NatsConnection
}

func (ue *UserEventClient) Connection(nc *pubsubinfra.NatsConnection) {
	ue.natConnection = nc
}

func (ue *UserEventClient) GetConnection() *pubsubinfra.NatsConnection {
	return ue.natConnection
}

func (ue *UserEventClient) GetEventPayload(user ua.User) ([]byte, error) {
	b, err := json.Marshal(user.GetAggregateEventPayload())
	if err != nil {
		return []byte{}, nil
	}

	return b, nil
}

// UserCreated publishes the user aggregate after successful user creation
func (uv *UserEventClient) UserCreated(ctx context.Context, user ua.User) error {
	pb := pubsubinfra.NewPublication[ua.User](uv.natConnection)

	j, err := uv.GetEventPayload(user)

	if err != nil {
		return err
	}

	return pb.Publish(UserCreatedEventSubject, j)
}
