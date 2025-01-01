package event

import (
	"encoding/json"

	pubsub "github.com/wizact/go-todo-api/pkg/event-library/pubsub"

	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

const UserDomainTopicName = "User"
const NewUserRegisteredEventSubjectName = "NewUserRegistered"

type UserEventClient struct {
	natConnection *pubsub.NatsConnection
}

func (ue *UserEventClient) Connection(nc *pubsub.NatsConnection) {
	ue.natConnection = nc
}

func (ue *UserEventClient) GetConnection() *pubsub.NatsConnection {
	return ue.natConnection
}

func (ue *UserEventClient) MarshalEventPayload(userDE ude.UserDomainEvent) ([]byte, error) {
	b, err := json.Marshal(userDE)
	if err != nil {
		return []byte{}, nil
	}

	return b, nil
}

// NewUserRegisteredEventFQN retuens the fully qualified name (FQN) for new user registered event
func (ue *UserEventClient) NewUserRegisteredEventFQN() string {
	return UserDomainTopicName + "." + NewUserRegisteredEventSubjectName
}
