package pubsub

type NatsClient[T any, A any] interface {
	*T
	Connection(*NatsConnection)
	GetConnection() *NatsConnection
	GetEventPayload(A) ([]byte, error)
}

type NatsClientFactory[T any, A any, P NatsClient[T, A]] struct {
	connection *NatsConnection
}

func (r *NatsClientFactory[T, A, P]) Get() (P, error) {
	if r.connection == nil {
		sc, err := NewNatsConnection("", "")

		if err != nil {
			return nil, err
		}
		r.connection = sc
	}

	var result P = new(T)

	result.Connection(r.connection)

	return result, nil
}
