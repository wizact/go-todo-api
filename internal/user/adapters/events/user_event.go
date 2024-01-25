package events

import (
	"context"
	"log"

	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserEvent struct {
	pubsub *pubsubinfra.PubSubConnection
}

func (uv *UserEvent) UserCreated(ctx context.Context, user ua.User) {
	uv.pubsub.Connect()
	log.Println(user)
}
