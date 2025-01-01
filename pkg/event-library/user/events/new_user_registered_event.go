package event

import (
	"context"

	"github.com/nats-io/nats.go"
	pubsub_infra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

// PublishNewUserRegisteredEvent publishes the user aggregate after successful user creation
func (uv *UserEventClient) PublishNewUserRegisteredEvent(ctx context.Context, userDE ude.UserDomainEvent) error {
	pb := pubsub_infra.NewPublication[ude.UserDomainEvent](uv.natConnection)

	j, err := uv.MarshalEventPayload(userDE)

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
