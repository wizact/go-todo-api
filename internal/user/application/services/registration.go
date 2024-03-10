package service

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/wizact/go-todo-api/internal/infra/pubsub"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	event_port "github.com/wizact/go-todo-api/internal/user/ports/events"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
)

// Registration application service responsible for managing the lifecycle of a user registration
type Registration struct {
	userEventClient event_port.UserEventClient
	// app service can reference domain service ( but not the other way arround)
	userAccountUseCase usecase_port.UserAccountUseCase
	done               chan bool
}

// NewRegisteration returns a new instance of Registration application service
func NewRegisteration(uec event_port.UserEventClient, uc usecase_port.UserAccountUseCase) *Registration {
	return &Registration{
		userEventClient:    uec,
		userAccountUseCase: uc,
		done:               make(chan bool),
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

			// Unmarshal domain event to get the (aggregate id)
			ude, e := r.getUserFromPayload(newUser.Data)
			if e != nil {
				log.Println("user registration app service > send email confirmation: ", e)
				continue
			}

			u, err := r.userAccountUseCase.GetUserById(context.Background(), ude.ID)
			if err != nil {
				log.Println("user registration app service > send email confirmation: ", err)
				continue
			}
			log.Println(u)
			//TODO: Load  the aggregate, and retrieve the verification token / salt
			/*
				1. get the token for the user
				2. create a jwt
				3. form the email dto
				4. trigger the email send
			*/
		case <-done:
			log.Println("unsubscribing NewUserRegisteredListener")
			unsubcb()
			break L
		}
	}
	return nil
}

func (r *Registration) getUserFromPayload(p []byte) (ua.UserDomainEvent, error) {
	u := ua.NewUser()
	return u.LoadDomainEventObject(p)

}
