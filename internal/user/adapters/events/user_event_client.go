package event

import (
	"encoding/json"

	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

const UserDomainTopicName = "User"
const NewUserRegisteredEventSubjectName = "NewUserRegistered"

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
	b, err := json.Marshal(user.GetDomainEventPayload())
	if err != nil {
		return []byte{}, nil
	}

	return b, nil
}

// NewUserRegisteredEventFQN retuens the fully qualified name (FQN) for new user registered event
func (ue *UserEventClient) NewUserRegisteredEventFQN() string {
	return UserDomainTopicName + "." + NewUserRegisteredEventSubjectName
}
