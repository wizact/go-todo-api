package user

import (
	"context"
	"log"

	userAppServicePort "github.com/wizact/go-todo-api/internal/user/ports/service"
	communicationPort "github.com/wizact/go-todo-api/pkg/communication/ports/service"
	eventInput "github.com/wizact/go-todo-api/pkg/eventlibrary/ports/input/event"
	de "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
)

// RegisteredEventListener application service responsible for managing the lifecycle of a user registration
type RegisteredEventListener struct {
	emailClient         communicationPort.Emailer
	userEventClient     eventInput.UserSubscriber
	registrationService userAppServicePort.Registration
	templateID          string
	done                chan bool
}

// NewNewUserRegisteredEventListene returns a new instance of RegisteredEventListener application service
func NewRegisteredEventListener(
	uec eventInput.UserSubscriber,
	registrationService userAppServicePort.Registration,
	emailClient communicationPort.Emailer,
	templateID string,
) *RegisteredEventListener {
	return &RegisteredEventListener{
		emailClient:         emailClient,
		userEventClient:     uec,
		registrationService: registrationService,
		templateID:          templateID,
		done:                make(chan bool),
	}
}

func (r *RegisteredEventListener) Done() {
	r.done <- true
}

// Listen listens to the event and trigger the lifecycle required for user approval process
func (r *RegisteredEventListener) Listen(ctx context.Context) error {
	nuc := make(chan de.UserDomainEvent)

	unsubcb, err := r.userEventClient.SubscribeToNewUserRegisteredEvent(ctx, nuc)

	if err != nil {
		return err
	}

	go r.sendUserEmailVerificationMessage(ctx, nuc, r.done, unsubcb)

	return nil
}

func (r *RegisteredEventListener) sendUserEmailVerificationMessage(
	ctx context.Context,
	nuc <-chan de.UserDomainEvent,
	done chan bool,
	unsubcb eventInput.Unsubscribe,
) error {
L:
	for {
		select {
		case ude, ok := <-nuc:
			if !ok {
				log.Println("communication > terminating NewUserRegisteredListener")
				unsubcb()
				break L
			}

			log.Println("communication > Preparing email verification message for:", ude.Email)

			// Call user app service to get the user info required to send the email
			ed, err := r.registrationService.FetchRegistrationVerificationEmailData(ctx, ude.ID)
			if err != nil {
				log.Println("communication > new user registered event listener app service > send email confirmation: ", err)
				continue
			}

			// Send the email
			if err := r.emailClient.SendUsingTemplate(ude.ID.String(), ude.Email, "User Registration Verification", r.templateID, ed); err != nil {
				log.Println("communication > send registration verification email:", err)
			}

		case <-done:
			log.Println("communication > unsubscribing NewUserRegisteredListener")
			unsubcb()
			break L
		case <-ctx.Done():
			log.Println("communication > cancelling NewUserRegisteredListener")
			unsubcb()
			break L
		}
	}
	return nil
}
