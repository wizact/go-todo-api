package pubsub

import (
	"errors"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNewNatsConnection_ExplicitOptionsUseArguments(t *testing.T) {
	t.Setenv("TODOAPI_NATSURL", "nats://environment:4222")

	connection, err := NewNatsConnection("nats://explicit:4222", "explicit-client")
	if err != nil {
		t.Fatalf("NewNatsConnection() error = %v", err)
	}

	got := *connection.options
	want := (PubSubOpt{Urls: "nats://explicit:4222", ClientName: "explicit-client"})
	if got != want {
		t.Fatalf("NewNatsConnection() options = %#v, want %#v", got, want)
	}
}

func TestNatsConnection_ConnectFailurePreservesCause(t *testing.T) {
	t.Parallel()

	cause := errors.New("dial failed")
	connection := &NatsConnection{
		options: &PubSubOpt{Urls: "nats://unavailable:4222", ClientName: "test-client"},
		connect: func(string, ...nats.Option) (*nats.Conn, error) {
			return nil, cause
		},
	}

	_, err := connection.Connect()

	if !errors.Is(err, ErrFailedToConnectToNats) || !errors.Is(err, cause) {
		t.Fatalf("Connect() error = %v, want connection sentinel and dial cause", err)
	}
}
