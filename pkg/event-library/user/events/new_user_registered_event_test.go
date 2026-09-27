package event

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

func TestForwardNewUserRegisteredEvent(t *testing.T) {
	t.Parallel()

	want := ude.UserDomainEvent{
		ID:               uuid.New(),
		FirstName:        "Ada",
		LastName:         "Lovelace",
		Email:            "ada@example.com",
		IsActive:         true,
		HasVerifiedEmail: true,
	}
	payload, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	events := make(chan ude.UserDomainEvent, 1)
	client := &UserEventClient{}

	if err := client.forwardNewUserRegisteredEvent(context.Background(), events, &nats.Msg{Data: payload}); err != nil {
		t.Fatalf("forward event: %v", err)
	}

	if got := <-events; got != want {
		t.Fatalf("forwarded event = %#v, want %#v", got, want)
	}
}

func TestForwardNewUserRegisteredEventRejectsInvalidPayload(t *testing.T) {
	t.Parallel()

	events := make(chan ude.UserDomainEvent, 1)
	client := &UserEventClient{}

	err := client.forwardNewUserRegisteredEvent(context.Background(), events, &nats.Msg{Data: []byte("not-json")})
	if err == nil {
		t.Fatal("forward event error = nil, want invalid payload error")
	}

	select {
	case got := <-events:
		t.Fatalf("forwarded unexpected event: %#v", got)
	default:
	}
}
