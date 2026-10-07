package usereventlistener

import (
	"context"
	"log"

	user_app_svc_port "github.com/wizact/go-todo-api/internal/user/ports/applications"
	comms_app_svc_port "github.com/wizact/go-todo-api/pkg/communication/ports/applications"
	event_input "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	de "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

// NewUserRegisteredEventListener application service responsible for managing the lifecycle of a user registration
type NewUserRegisteredEventListener struct {
	emailClientAppSvc comms_app_svc_port.Emailer
	userEventClient   event_input.UserEventClientInput
	userRegAppSvc     user_app_svc_port.Registration
	templateID        string
	done              chan bool
}

// NewNewUserRegisteredEventListene returns a new instance of NewUserRegisteredEventListener application service
func NewNewUserRegisteredEventListener(
	uec event_input.UserEventClientInput,
	userRegAppSvc user_app_svc_port.Registration,
	emailClientAppSvc comms_app_svc_port.Emailer,
	templateID string,
) *NewUserRegisteredEventListener {
	return &NewUserRegisteredEventListener{
		emailClientAppSvc: emailClientAppSvc,
		userEventClient:   uec,
		userRegAppSvc:     userRegAppSvc,
		templateID:        templateID,
		done:              make(chan bool),
	}
}

func (r *NewUserRegisteredEventListener) Done() {
	r.done <- true
}

// Listen listens to the event and trigger the lifecycle required for user approval process
func (r *NewUserRegisteredEventListener) Listen(ctx context.Context) error {
	nuc := make(chan de.UserDomainEvent)

	unsubcb, err := r.userEventClient.SubscribeToNewUserRegisteredEvent(ctx, nuc)

	if err != nil {
		return err
	}

	go r.sendUserEmailVerificationMessage(ctx, nuc, r.done, unsubcb)

	return nil
}

func (r *NewUserRegisteredEventListener) sendUserEmailVerificationMessage(
	ctx context.Context,
	nuc <-chan de.UserDomainEvent,
	done chan bool,
	unsubcb event_input.Unsubscribe,
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
			ed, err := r.userRegAppSvc.GetRegistrationVerificationEmailData(ctx, ude.ID)
			if err != nil {
				log.Println("communication > new user registered event listener app service > send email confirmation: ", err)
				continue
			}

			// Send the email
			if err := r.emailClientAppSvc.SendUsingTemplate(ude.ID.String(), ude.Email, "User Registration Verification", r.templateID, ed); err != nil {
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
