package event

import (
	"context"

	"github.com/nats-io/nats.go"
	pubsub_infra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

// PublishNewUserRegisteredEvent publishes the user aggregate after successful user creation
func (uv *UserEventClient) PublishNewUserRegisteredEvent(ctx context.Context, user ua.User) error {
	pb := pubsub_infra.NewPublication[ua.User](uv.natConnection)

	j, err := uv.MarshalEventPayload(user)

	if err != nil {
		return err
	}

	return pb.Publish(uv.NewUserRegisteredEventFQN(), j)
}

func (uv *UserEventClient) SubscribeToNewUserRegisteredEvent(ctx context.Context, ch chan *nats.Msg) (pubsub_infra.ChannelUnsubscribeCallBack, error) {
	sub := pubsub_infra.NewSubscription(uv.natConnection)
	unsubcf, err := sub.SubscribeChan(uv.NewUserRegisteredEventFQN(), ch)
	if err != nil {
		return nil, err
	}

	return unsubcf, nil
}
