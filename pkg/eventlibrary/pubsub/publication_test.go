package pubsub

import (
	"errors"
	"strings"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestPublication_PublishSetConnectionFailureAddsContext(t *testing.T) {
	t.Parallel()

	cause := errors.New("dial failed")
	connection := &NATSConnection{
		options: &Options{URLs: "nats://unavailable:4222", ClientName: "test-client"},
		connect: func(string, ...nats.Option) (*nats.Conn, error) {
			return nil, cause
		},
	}
	publication := NewPublication[struct{}](connection)

	err := publication.Publish("test.subject", []byte("payload"))

	if !errors.Is(err, ErrConnectNATS) || !errors.Is(err, cause) || !strings.Contains(err.Error(), "connect publisher") {
		t.Fatalf("Publish() error = %v, want publisher context and connection causes", err)
	}
}
