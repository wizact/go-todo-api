package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	event_input "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
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

func (uv *UserEventClient) SubscribeToNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent) (event_input.Unsubscribe, error) {
	sub := pubsub_infra.NewSubscription(uv.natConnection)
	err := sub.Subscribe(uv.NewUserRegisteredEventFQN(), func(message *nats.Msg) {
		if err := uv.forwardNewUserRegisteredEvent(ctx, events, message); err != nil {
			log.Printf("forward new user registered event: %v", err)
		}
	})
	if err != nil {
		return nil, err
	}

	return event_input.Unsubscribe(sub.UnsubscribeFn()), nil
}

func (uv *UserEventClient) forwardNewUserRegisteredEvent(ctx context.Context, events chan<- ude.UserDomainEvent, message *nats.Msg) error {
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
