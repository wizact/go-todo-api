package pubsub

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	ErrResolveNATSOptions = errors.New("failed to resolve nats options")
	ErrConnectNATS        = errors.New("failed to connect to nats")
)

type Options struct {
	URLs       string
	ClientName string
}

type NATSConnection struct {
	conn    *nats.Conn
	options *Options
	connect func(string, ...nats.Option) (*nats.Conn, error)
}

// NewNATSConnection creates a NATS connection configuration for the supplied
// cluster URLs and client name. The connection is established lazily.
func NewNATSConnection(urls, clientName string) (*NATSConnection, error) {
	nc, err := resolveConnectionOpts(urls, clientName)

	if err != nil {
		return nil, err
	}

	return &NATSConnection{options: nc, connect: nats.Connect}, nil
}

func resolveConnectionOpts(urls, clientName string) (*Options, error) {
	if urls != "" && clientName != "" {
		return &Options{URLs: urls, ClientName: clientName}, nil
	}

	pso := Config{}
	nu, cn, err := pso.Resolve()

	if err != nil {
		return nil, ErrResolveNATSOptions
	}

	return &Options{URLs: nu, ClientName: cn}, nil
}

// Connect returns the active connection or establishes it on first use.
func (n *NATSConnection) Connect() (*nats.Conn, error) {
	if n.conn != nil && n.conn.IsConnected() {
		return n.conn, nil
	}

	opts := []nats.Option{nats.Name(n.options.ClientName)}
	opts = setupConnOptions(opts)

	nc, err := n.connect(n.options.URLs, opts...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrConnectNATS, err)
	}

	n.conn = nc

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
