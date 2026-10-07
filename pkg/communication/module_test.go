package communication

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	userportmocks "github.com/wizact/go-todo-api/internal/user/ports/mocks"
	eventinput "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	userdomain "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

func TestInstantiateUserDomainListenersAndListen_UsesProvidedRegistration(t *testing.T) {
	controller := gomock.NewController(t)
	registration := userportmocks.NewMockRegistration(controller)
	userID := uuid.New()
	registrationCalled := make(chan struct{})
	registration.EXPECT().
		GetRegistrationVerificationEmailData(gomock.Any(), userID).
		DoAndReturn(func(context.Context, uuid.UUID) (map[string]string, error) {
			close(registrationCalled)
			return map[string]string{"token": "verification-token"}, nil
		})

	subscriber := &moduleUserEventSubscriberStub{}
	listener := instantiateUserDomainListenersAndListen(subscriber, registration, moduleEmailerStub{})
	t.Cleanup(listener.Done)
	subscriber.events <- userdomain.UserDomainEvent{
		ID:    userID,
		Email: "ada@example.com",
	}

	select {
	case <-registrationCalled:
	case <-time.After(time.Second):
		t.Fatal("provided registration service was not called")
	}
}

type moduleUserEventSubscriberStub struct {
	events chan<- userdomain.UserDomainEvent
}

func (subscriber *moduleUserEventSubscriberStub) SubscribeToNewUserRegisteredEvent(
	_ context.Context,
	events chan<- userdomain.UserDomainEvent,
) (eventinput.Unsubscribe, error) {
	subscriber.events = events
	return func() error { return nil }, nil
}

type moduleEmailerStub struct{}

func (moduleEmailerStub) Send(string, string, string, string, string) error {
	return nil
}

func (moduleEmailerStub) SendUsingTemplate(string, string, string, string, map[string]string) error {
	return nil
}
