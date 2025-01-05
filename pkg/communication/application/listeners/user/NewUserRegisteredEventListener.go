package usereventlistener

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	user_app_svc_port "github.com/wizact/go-todo-api/internal/user/ports/applications"
	comms_app_svc_port "github.com/wizact/go-todo-api/pkg/communication/ports/applications"
	event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/events"
	"github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	de "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

// TODO: Replace with actual template id
const VERIFY_REGISTRATION_TEMPLATE_ID = ""

// NewUserRegisteredEventListener application service responsible for managing the lifecycle of a user registration
type NewUserRegisteredEventListener struct {
	emailClientAppSvc comms_app_svc_port.Emailer
	userEventClient   event_port.UserEventClient
	userRegAppSvc     user_app_svc_port.Registration
	done              chan bool
}

// NewNewUserRegisteredEventListene returns a new instance of NewUserRegisteredEventListener application service
func NewNewUserRegisteredEventListener(uec event_port.UserEventClient, userRegAppSvc user_app_svc_port.Registration, emailClientAppSvc comms_app_svc_port.Emailer) *NewUserRegisteredEventListener {
	return &NewUserRegisteredEventListener{
		emailClientAppSvc: emailClientAppSvc,
		userEventClient:   uec,
		userRegAppSvc:     userRegAppSvc,
		done:              make(chan bool),
	}
}

func (r *NewUserRegisteredEventListener) Done() {
	r.done <- true
}

// Listen listens to the event and trigger the lifecycle required for user approval process
func (r *NewUserRegisteredEventListener) Listen() error {
	nuc := make(chan *nats.Msg)

	unsubcb, err := r.userEventClient.SubscribeToNewUserRegisteredEvent(context.Background(), nuc)

	if err != nil {
		return err
	}

	go r.sendUserEmailVerificationMessage(nuc, r.done, unsubcb)

	return nil
}

func (r *NewUserRegisteredEventListener) sendUserEmailVerificationMessage(nuc <-chan *nats.Msg, done chan bool, unsubcb pubsub.ChannelUnsubscribeCallBack) error {
L:
	for {
		select {
		case newUser, ok := <-nuc:
			if !ok {
				log.Println("communication > terminating NewUserRegisteredListener")
				unsubcb()
				break L
			}

			// Unmarshal domain event to get the (aggregate id)
			ude, e := r.getUserFromPayload(newUser.Data)
			if e != nil {
				log.Println("communication > new user registered event listener app service > send email confirmation: ", e)
				continue
			}

			log.Println("communication > Preparing email verification message for:", ude.Email)

			// Call user app service to get the user info required to send the email
			ed, err := r.userRegAppSvc.GetRegistrationVerificationEmailData(ude.ID)
			if err != nil {
				log.Println("communication > new user registered event listener app service > send email confirmation: ", err)
			}

			// Send the email
			r.emailClientAppSvc.SendUsingTemplate(ude.ID.String(), ude.Email, "User Registration Verification", VERIFY_REGISTRATION_TEMPLATE_ID, ed)

		case <-done:
			log.Println("communication > unsubscribing NewUserRegisteredListener")
			unsubcb()
			break L
		}
	}
	return nil
}

func (r *NewUserRegisteredEventListener) getUserFromPayload(p []byte) (de.UserDomainEvent, error) {
	ude := de.UserDomainEvent{}
	if err := ude.LoadDomainEventObject(p); err != nil {
		return ude, err
	}

	return ude, nil

}
