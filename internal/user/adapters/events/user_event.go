package events

import (
	"context"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	"log"
)

type UserEvent struct {
}

func (uv *UserEvent) UserCreated(ctx context.Context, user ua.User) {
	log.Println(user)
}
