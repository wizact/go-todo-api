package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	userAppServiceMocks "github.com/wizact/go-todo-api/internal/user/ports/mocks"
	userEventInput "github.com/wizact/go-todo-api/pkg/eventlibrary/ports/input/event"
	ude "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
	"go.uber.org/mock/gomock"
)

func TestRegisteredEventListenerSendsVerificationEmail(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	registration := userAppServiceMocks.NewMockRegistration(controller)
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
		FetchRegistrationVerificationEmailData(listenerContext, userID).
		Return(templateData, nil)

	listener := NewRegisteredEventListener(subscriber, registration, emailer, "verification-template")
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
		if got.templateID != "verification-template" {
			t.Errorf("email template ID = %q, want %q", got.templateID, "verification-template")
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

func TestRegisteredEventListener_CancelledContextUnsubscribes(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	registration := userAppServiceMocks.NewMockRegistration(controller)
	subscriber := &userEventSubscriberStub{unsubscribed: make(chan struct{})}
	emailer := &emailerSpy{calls: make(chan templateEmail, 1)}
	listenerContext, cancel := context.WithCancel(context.Background())
	listener := NewRegisteredEventListener(subscriber, registration, emailer, "verification-template")
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

func TestRegisteredEventListener_RegistrationFailureSkipsEmail(t *testing.T) {
	t.Parallel()

	controller := gomock.NewController(t)
	registration := userAppServiceMocks.NewMockRegistration(controller)
	subscriber := &userEventSubscriberStub{unsubscribed: make(chan struct{})}
	emailer := &emailerSpy{calls: make(chan templateEmail, 1)}
	registrationCalled := make(chan struct{})
	listenerContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	userID := uuid.New()

	registration.EXPECT().
		FetchRegistrationVerificationEmailData(listenerContext, userID).
		DoAndReturn(func(context.Context, uuid.UUID) (map[string]string, error) {
			close(registrationCalled)
			return nil, errors.New("issue credential")
		})

	listener := NewRegisteredEventListener(subscriber, registration, emailer, "verification-template")
	if err := listener.Listen(listenerContext); err != nil {
		t.Fatalf("listen: %v", err)
	}
	subscriber.events <- ude.UserDomainEvent{ID: userID, Email: "ada@example.com"}

	select {
	case <-registrationCalled:
	case <-time.After(time.Second):
		t.Fatal("registration service was not called")
	}

	select {
	case email := <-emailer.calls:
		t.Fatalf("unexpected verification email: %#v", email)
	case <-time.After(100 * time.Millisecond):
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

func (s *userEventSubscriberStub) SubscribeToNewUserRegisteredEvent(_ context.Context, events chan<- ude.UserDomainEvent) (userEventInput.Unsubscribe, error) {
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
	templateID   string
	templateData map[string]string
}

type emailerSpy struct {
	calls chan templateEmail
}

func (e *emailerSpy) Send(_, _, _, _, _ string) error {
	return nil
}

func (e *emailerSpy) SendUsingTemplate(to, email, subject, templateID string, templateData map[string]string) error {
	e.calls <- templateEmail{
		to:           to,
		email:        email,
		subject:      subject,
		templateID:   templateID,
		templateData: templateData,
	}
	return nil
}
