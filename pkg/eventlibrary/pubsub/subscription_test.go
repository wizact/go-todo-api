package pubsub

import (
	"errors"
	"strings"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestSubscription_SubscribeSetConnectionFailureAddsContext(t *testing.T) {
	t.Parallel()

	cause := errors.New("dial failed")
	connection := &NATSConnection{
		options: &Options{URLs: "nats://unavailable:4222", ClientName: "test-client"},
		connect: func(string, ...nats.Option) (*nats.Conn, error) {
			return nil, cause
		},
	}
	subscription := NewSubscription(connection)

	err := subscription.Subscribe("test.subject", func(*nats.Msg) {})

	if !errors.Is(err, ErrConnectNATS) || !errors.Is(err, cause) || !strings.Contains(err.Error(), "connect subscriber") {
		t.Fatalf("Subscribe() error = %v, want subscriber context and connection causes", err)
	}
}
