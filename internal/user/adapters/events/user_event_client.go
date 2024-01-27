package event

import (
	"encoding/json"

	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

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
