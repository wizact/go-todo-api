package pubsub

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type Subscription struct {
	psc          *NATSConnection
	subscription *nats.Subscription
}

func NewSubscription(psc *NATSConnection) Subscription {
	return Subscription{psc: psc}
}

func (s *Subscription) connect() error {
	_, err := s.psc.Connect()

	if err != nil {
		return fmt.Errorf("connect subscriber: %w", err)
	}
	return nil
}

func (s *Subscription) Subscribe(subj string, sc nats.MsgHandler) error {
	if err := s.connect(); err != nil {
		return err
	}

	sub, err := s.psc.conn.Subscribe(subj, sc)

	if err != nil {
		return err
	}

	s.subscription = sub

	return nil
}

func (s *Subscription) SubscribeChan(subj string, ch chan *nats.Msg) (UnsubscribeFunc, error) {
	if err := s.connect(); err != nil {
		return nil, err
	}

	sub, err := s.psc.conn.ChanSubscribe(subj, ch)

	if err != nil {
		return nil, err
	}

	s.subscription = sub

	return s.UnsubscribeFunc(), nil
}

func (s *Subscription) UnsubscribeFunc() UnsubscribeFunc {
	return func() error {
		s.psc.conn.Flush()
		return s.subscription.Unsubscribe()
	}
}

type UnsubscribeFunc func() error
