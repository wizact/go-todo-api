package pubsub

type NATSClient[T any, A any] interface {
	*T
	SetConnection(*NATSConnection)
	Connection() *NATSConnection
	MarshalEventPayload(A) ([]byte, error)
}

type NATSClientFactory[T any, A any, P NATSClient[T, A]] struct {
	connection *NATSConnection
}

func (f *NATSClientFactory[T, A, P]) Create() (P, error) {
	if f.connection == nil {
		sc, err := NewNATSConnection("", "")

		if err != nil {
			return nil, err
		}
		f.connection = sc
	}

	var result P = new(T)

	result.SetConnection(f.connection)

	return result, nil
}
