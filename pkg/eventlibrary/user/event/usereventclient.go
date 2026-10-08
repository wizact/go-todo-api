package event

import (
	"encoding/json"

	pubsub "github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"

	ude "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
)

const UserDomainTopicName = "User"
const NewUserRegisteredEventSubjectName = "NewUserRegistered"

type UserEventClient struct {
	natsConnection *pubsub.NATSConnection
}

func (ue *UserEventClient) SetConnection(nc *pubsub.NATSConnection) {
	ue.natsConnection = nc
}

func (ue *UserEventClient) Connection() *pubsub.NATSConnection {
	return ue.natsConnection
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
