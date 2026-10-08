package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	eventInput "github.com/wizact/go-todo-api/pkg/eventlibrary/ports/input/event"
	pubsubInfra "github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	ude "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
)

// PublishNewUserRegisteredEvent publishes the user aggregate after successful user creation
func (ue *UserEventClient) PublishNewUserRegisteredEvent(ctx context.Context, userDE ude.UserDomainEvent) error {
	pb := pubsubInfra.NewPublication[ude.UserDomainEvent](ue.natsConnection)

	j, err := ue.MarshalEventPayload(userDE)

	if err != nil {
		return err
	}

	return pb.Publish(ue.NewUserRegisteredEventFQN(), j)
}

func (ue *UserEventClient) SubscribeToNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent) (eventInput.Unsubscribe, error) {
	sub := pubsubInfra.NewSubscription(ue.natsConnection)
	err := sub.Subscribe(ue.NewUserRegisteredEventFQN(), func(message *nats.Msg) {
		if err := ue.forwardNewUserRegisteredEvent(ctx, events, message); err != nil {
			log.Printf("forward new user registered event: %v", err)
		}
	})
	if err != nil {
		return nil, err
	}

	return eventInput.Unsubscribe(sub.UnsubscribeFunc()), nil
}

func (ue *UserEventClient) forwardNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent, message *nats.Msg) error {
	userEvent := ude.UserDomainEvent{}
	if err := json.Unmarshal(message.Data, &userEvent); err != nil {
		return fmt.Errorf("decode event payload: %w", err)
	}

	select {
	case events <- userEvent:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("forward event: %w", ctx.Err())
	}
}
