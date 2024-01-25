package pubsub

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/internal/infra"
)

type NatsConfig struct {
	NatsUrl string
}

// GetUrl gets the path to message queue from the env variable and client name
func (e *NatsConfig) GetConfig() (string, string, error) {
	if e.NatsUrl != "" {
		return e.NatsUrl, infra.APPNAME, nil
	}

	err := envconfig.Process(infra.APPNAME, e)
	if err != nil {
		panic(err)
	}

	if e.NatsUrl == "" {
		return "", infra.APPNAME, errors.New("cannot resolve nats url")
	}

	return e.NatsUrl, infra.APPNAME, nil
}
