package pubsub

import (
	"errors"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	ErrFailedToResolveDbNatsOpt = errors.New("failed to resolve nats options")
	ErrFailedToConnectToNats    = errors.New("failed to connect to nats")
)

type PubSubOpt struct {
	Urls       string
	ClientName string
}

type NatsConnection struct {
	conn    *nats.Conn
	options *PubSubOpt
}

// NewNatsConnection creates and returns a new nats connection with urls of the cluster and client name.
// The connection is not established at this stage.
// Urls are a string a instances of cluster seperated by space, or a single url.
// If urls and client name are not provided, they will be resolved from env variables.
func NewNatsConnection(urls, clientName string) (*NatsConnection, error) {
	nc, err := resolveConnectionOpts("", "")

	if err != nil {
		return nil, err
	}

	return &NatsConnection{options: nc}, nil
}

func resolveConnectionOpts(urls, clientName string) (*PubSubOpt, error) {
	if urls != "" && clientName != "" {
		return &PubSubOpt{Urls: urls, ClientName: clientName}, nil
	}

	pso := NatsConfig{}
	nu, cn, err := pso.GetConfig()

	if err != nil {
		return nil, ErrFailedToResolveDbNatsOpt
	}

	return &PubSubOpt{Urls: nu, ClientName: cn}, nil
}

// Connect returns an already established connection or establish a new connection and returns it.
func (psc *NatsConnection) Connect() (*nats.Conn, error) {
	if psc.conn != nil && psc.conn.IsConnected() {
		return psc.conn, nil
	}

	opts := []nats.Option{nats.Name(psc.options.ClientName)}
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(psc.options.Urls, opts...)
	if err != nil {
		return nil, ErrFailedToConnectToNats
	}

	psc.conn = nc

	return nc, nil
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
