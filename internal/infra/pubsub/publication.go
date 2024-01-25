package pubsub

import (
	"encoding/json"
)

type Publication[T any] struct {
	psc *NatsConnection
}

func NewPublication[T any](conn *NatsConnection) Publication[T] {
	return Publication[T]{psc: conn}
}

func (p *Publication[T]) connect() error {
	_, err := p.psc.Connect()

	if err != nil {
		return err
	}
	return nil
}

func (p *Publication[T]) Publish(subj string, data []byte) error {

	if err := p.connect(); err != nil {
		return err
	}

	return p.psc.conn.Publish(subj, data)
}

func (p *Publication[T]) PublishObject(subj string, o T) error {
	if err := p.connect(); err != nil {
		return err
	}

	jo, err := json.Marshal(o)

	if err != nil {
		return err
	}

	return p.psc.conn.Publish(subj, jo)
}
