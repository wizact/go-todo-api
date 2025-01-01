package pubsub

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/pkg/version"
)

type NatsConfig struct {
	NatsUrl string
}

// GetUrl gets the path to message queue from the env variable and client name
func (e *NatsConfig) GetConfig() (string, string, error) {
	if e.NatsUrl != "" {
		return e.NatsUrl, version.APPNAME, nil
	}

	err := envconfig.Process(version.APPNAME, e)
	if err != nil {
		panic(err)
	}

	if e.NatsUrl == "" {
		return "", version.APPNAME, errors.New("cannot resolve nats url")
	}

	return e.NatsUrl, version.APPNAME, nil
}
