package pubsub

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type PubSubOpt struct {
	Urls       string
	ClientName string
}

type PubSub struct {
	conn *nats.Conn
}

func (ps *PubSub) Connect(pbo PubSubOpt) *nats.Conn {
	if ps.conn != nil {
		return ps.conn
	}

	opts := []nats.Option{nats.Name(pbo.ClientName)}
	opts = setupConnOptions(opts)

	nc, err := nats.Connect(pbo.Urls, opts...)
	if err != nil {
		log.Fatal(err)
	}

	ps.conn = nc

	return nc
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
