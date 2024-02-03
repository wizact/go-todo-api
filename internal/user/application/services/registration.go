package service

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/wizact/go-todo-api/internal/infra/pubsub"
	event_port "github.com/wizact/go-todo-api/internal/user/ports/events"
)

// Registration application service responsible for managing the lifecycle of a user registration
type Registration struct {
	userEventClient event_port.UserEventClient
	done            chan bool
}

// NewRegisteration returns a new instance of Registration application service
func NewRegisteration(uec event_port.UserEventClient) *Registration {

	return &Registration{
		userEventClient: uec,
		done:            make(chan bool),
	}
}

func (r *Registration) Done() {
	r.done <- true
}

// NewUserRegisteredListener listens to the event and trigger the lifecycle required for user approval process
func (r *Registration) NewUserRegisteredListener() error {
	nuc := make(chan *nats.Msg)

	unsubcb, err := r.userEventClient.SubscribeToNewUserRegisteredEvent(context.Background(), nuc)

	if err != nil {
		return err
	}

	go r.sendUserEmailConfirmationMessage(nuc, r.done, unsubcb)

	return nil
}

func (r *Registration) sendUserEmailConfirmationMessage(nuc <-chan *nats.Msg, done chan bool, unsubcb pubsub.ChannelUnsubscribeCallBack) error {
L:
	for {
		select {
		case newUser, ok := <-nuc:
			if !ok {
				log.Println("terminating NewUserRegisteredListener")
				unsubcb()
				break L
			}

			log.Printf("%v, %v \n", newUser.Subject, string(newUser.Data))
		case <-done:
			log.Println("unsubscribing NewUserRegisteredListener")
			unsubcb()
			break L
		}
	}
	return nil
}
