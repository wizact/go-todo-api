package pubsub

import (
	"errors"
	"strings"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestPublication_PublishConnectionFailureAddsContext(t *testing.T) {
	t.Parallel()

	cause := errors.New("dial failed")
	connection := &NatsConnection{
		options: &PubSubOpt{Urls: "nats://unavailable:4222", ClientName: "test-client"},
		connect: func(string, ...nats.Option) (*nats.Conn, error) {
			return nil, cause
		},
	}
	publication := NewPublication[struct{}](connection)

	err := publication.Publish("test.subject", []byte("payload"))

	if !errors.Is(err, ErrFailedToConnectToNats) || !errors.Is(err, cause) || !strings.Contains(err.Error(), "connect publisher") {
		t.Fatalf("Publish() error = %v, want publisher context and connection causes", err)
	}
}
