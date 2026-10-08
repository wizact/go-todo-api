package pubsub

import (
	"errors"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNewNATSSetConnection_ExplicitOptionsUseArguments(t *testing.T) {
	t.Setenv("TODOAPI_NATSURL", "nats://environment:4222")

	connection, err := NewNATSConnection("nats://explicit:4222", "explicit-client")
	if err != nil {
		t.Fatalf("NewNatsConnection() error = %v", err)
	}

	got := *connection.options
	want := (Options{URLs: "nats://explicit:4222", ClientName: "explicit-client"})
	if got != want {
		t.Fatalf("NewNatsConnection() options = %#v, want %#v", got, want)
	}
}

func TestNATSSetConnection_ConnectFailurePreservesCause(t *testing.T) {
	t.Parallel()

	cause := errors.New("dial failed")
	connection := &NATSConnection{
		options: &Options{URLs: "nats://unavailable:4222", ClientName: "test-client"},
		connect: func(string, ...nats.Option) (*nats.Conn, error) {
			return nil, cause
		},
	}

	_, err := connection.Connect()

	if !errors.Is(err, ErrConnectNATS) || !errors.Is(err, cause) {
		t.Fatalf("Connect() error = %v, want connection sentinel and dial cause", err)
	}
}
