package pubsub

import (
	"github.com/nats-io/nats.go"
)

type Subscription struct {
	psc          *NatsConnection
	subscription *nats.Subscription
}

func NewSubscription(psc *NatsConnection) Subscription {
	return Subscription{psc: psc}
}

func (s *Subscription) connect() error {
	_, err := s.psc.Connect()

	if err != nil {
		return err
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

func (s *Subscription) SubscribeChan(subj string, ch chan *nats.Msg) (ChannelUnsubscribeCallBack, error) {
	if err := s.connect(); err != nil {
		return nil, err
	}

	sub, err := s.psc.conn.ChanSubscribe(subj, ch)

	if err != nil {
		return nil, err
	}

	s.subscription = sub

	return s.UnsubscribeFn(), nil
}

func (s *Subscription) UnsubscribeFn() ChannelUnsubscribeCallBack {
	return func() error {
		s.psc.conn.Flush()
		return s.subscription.Unsubscribe()
	}
}

type ChannelUnsubscribeCallBack func() error
