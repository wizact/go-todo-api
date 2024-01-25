package pubsub

import "encoding/json"

type Publication[T any] struct {
	psc *PubSubConnection
}

func NewPublication[T any](conn *PubSubConnection) Publication[T] {
	return Publication[T]{psc: conn}
}

func (p *Publication[T]) Publish(subj string, data []byte) error {
	return p.psc.conn.Publish(subj, data)
}

func (p *Publication[T]) PublishObject(subj string, o T) error {
	jo, err := json.Marshal(o)

	if err != nil {
		return err
	}

	return p.psc.conn.Publish(subj, jo)
}
