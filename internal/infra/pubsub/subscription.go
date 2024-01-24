package pubsub

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Subscription struct {
	pb           *PubSub
	subscription *nats.Subscription
}

func NewSubscription(conn *PubSub) Subscription {
	return Subscription{pb: conn}
}

func (s *Subscription) Subscribe(subj string) {
	sub, e := s.pb.conn.Subscribe(subj, func(msg *nats.Msg) {
		log.Println("Recieved %v with %v", msg.Header, msg.Data)
	})

	if e != nil {
		log.Fatalf(e.Error(), subj)
	}

	s.subscription = sub
}

func (s *Subscription) SubscribeChan(subj string, ch chan *nats.Msg) {
	sub, e := s.pb.conn.ChanSubscribe(subj, ch)

	if e != nil {
		log.Fatalf(e.Error(), subj)
	}

	s.subscription = sub
}

func (s *Subscription) Unsubscribe() error {
	s.pb.conn.Flush()
	return s.subscription.Unsubscribe()
}
