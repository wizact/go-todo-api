package event

import (
	"context"

	pubsub_infra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

const NewUserRegisteredEventSubject = "NewUserRegistered"

// NewUserRegistered publishes the user aggregate after successful user creation
func (uv *UserEventClient) NewUserRegistered(ctx context.Context, user ua.User) error {
	pb := pubsub_infra.NewPublication[ua.User](uv.natConnection)

	j, err := uv.GetEventPayload(user)

	if err != nil {
		return err
	}

	return pb.Publish(NewUserRegisteredEventSubject, j)
}
