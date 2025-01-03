package event

import (
	"context"

	"github.com/nats-io/nats.go"
	pubsub_infra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
)

type UserEventClientInput interface {
	SubscribeToNewUserRegisteredEvent(ctx context.Context, ch chan *nats.Msg) (pubsub_infra.ChannelUnsubscribeCallBack, error)
}
