package pubsub

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subscription struct {
	psc          *NatsConnection
	subscription *nats.Subscription
}

func NewSubscription(psc *NatsConnection) Subscription {
	return Subscription{psc: psc}
}

func (s *Subscription) Subscribe(subj string) error {
	sub, err := s.psc.conn.Subscribe(subj, func(msg *nats.Msg) {
		log.Printf("Recieved %v with %v \n", msg.Header, msg.Data)
	})

	if err != nil {
		return err
	}

	s.subscription = sub

	return nil
}

func (s *Subscription) SubscribeChan(subj string, ch chan *nats.Msg) error {
	sub, err := s.psc.conn.ChanSubscribe(subj, ch)

	if err != nil {
		return err
	}

	s.subscription = sub

	return nil
}

func (s *Subscription) Unsubscribe() error {
	s.psc.conn.Flush()
	return s.subscription.Unsubscribe()
}
