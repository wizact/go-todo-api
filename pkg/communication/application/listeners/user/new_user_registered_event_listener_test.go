package usereventlistener

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	user_app_svc_mocks "github.com/wizact/go-todo-api/internal/user/ports/mocks"
	user_event_input "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	ude "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
	"go.uber.org/mock/gomock"
)

func TestNewUserRegisteredEventListenerSendsVerificationEmail(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	registration := user_app_svc_mocks.NewMockRegistration(controller)
	subscriber := &userEventSubscriberStub{unsubscribed: make(chan struct{})}
	emailer := &emailerSpy{calls: make(chan templateEmail, 1)}
	userID := uuid.New()
	listenerContext := context.WithValue(context.Background(), listenerContextKey{}, "registration")
	event := ude.UserDomainEvent{
		ID:        userID,
		Email:     "ada@example.com",
		FirstName: "Ada",
		LastName:  "Lovelace",
	}
	templateData := map[string]string{"hash": "verification-hash"}

	registration.EXPECT().
		GetRegistrationVerificationEmailData(listenerContext, userID).
		Return(templateData, nil)

	listener := NewNewUserRegisteredEventListener(subscriber, registration, emailer)
	if err := listener.Listen(listenerContext); err != nil {
		t.Fatalf("listen: %v", err)
	}

	subscriber.events <- event

	select {
	case got := <-emailer.calls:
		if got.to != userID.String() {
			t.Errorf("email recipient ID = %q, want %q", got.to, userID)
		}
		if got.email != event.Email {
			t.Errorf("email address = %q, want %q", got.email, event.Email)
		}
		if got.subject != "User Registration Verification" {
			t.Errorf("email subject = %q", got.subject)
		}
		if got.templateData["hash"] != templateData["hash"] {
			t.Errorf("email template hash = %q, want %q", got.templateData["hash"], templateData["hash"])
		}
	case <-time.After(time.Second):
		t.Fatal("verification email was not sent")
	}

	listener.Done()
	select {
	case <-subscriber.unsubscribed:
	case <-time.After(time.Second):
		t.Fatal("event subscription was not cancelled")
	}
}

func TestNewUserRegisteredEventListener_CancelledContextUnsubscribes(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	registration := user_app_svc_mocks.NewMockRegistration(controller)
	subscriber := &userEventSubscriberStub{unsubscribed: make(chan struct{})}
	emailer := &emailerSpy{calls: make(chan templateEmail, 1)}
	listenerContext, cancel := context.WithCancel(context.Background())
	listener := NewNewUserRegisteredEventListener(subscriber, registration, emailer)
	if err := listener.Listen(listenerContext); err != nil {
		t.Fatalf("listen: %v", err)
	}

	cancel()

	select {
	case <-subscriber.unsubscribed:
	case <-time.After(time.Second):
		t.Fatal("event subscription was not cancelled")
	}
}

type listenerContextKey struct{}

type userEventSubscriberStub struct {
	events       chan<- ude.UserDomainEvent
	unsubscribed chan struct{}
}

func (s *userEventSubscriberStub) SubscribeToNewUserRegisteredEvent(_ context.Context, events chan<- ude.UserDomainEvent) (user_event_input.Unsubscribe, error) {
	s.events = events
	return func() error {
		close(s.unsubscribed)
		return nil
	}, nil
}

type templateEmail struct {
	to           string
	email        string
	subject      string
	templateData map[string]string
}

type emailerSpy struct {
	calls chan templateEmail
}

func (e *emailerSpy) Send(_, _, _, _, _ string) (int, error) {
	return 200, nil
}

func (e *emailerSpy) SendUsingTemplate(to, email, subject, _ string, templateData map[string]string) (int, error) {
	e.calls <- templateEmail{
		to:           to,
		email:        email,
		subject:      subject,
		templateData: templateData,
	}
	return 200, nil
}
