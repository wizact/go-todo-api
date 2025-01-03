package usereventlistener

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/events"
	"github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	de "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
)

// NewUserRegisteredEventListener application service responsible for managing the lifecycle of a user registration
type NewUserRegisteredEventListener struct {
	userEventClient event_port.UserEventClient
	done            chan bool
}

// NewNewUserRegisteredEventListene returns a new instance of NewUserRegisteredEventListener application service
func NewNewUserRegisteredEventListener(uec event_port.UserEventClient) *NewUserRegisteredEventListener {
	return &NewUserRegisteredEventListener{
		userEventClient: uec,
		done:            make(chan bool),
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
				log.Println("terminating NewUserRegisteredListener")
				unsubcb()
				break L
			}

			// Unmarshal domain event to get the (aggregate id)
			ude, e := r.getUserFromPayload(newUser.Data)
			if e != nil {
				log.Println("communication > new user registered event listener app service > send email confirmation: ", e)
				continue
			}

			log.Println(ude.Email)

			// Call user app service to get the user info required to send the email
			// Send the email

		case <-done:
			log.Println("unsubscribing NewUserRegisteredListener")
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
