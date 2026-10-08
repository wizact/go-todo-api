package pubsub

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/pkg/version"
)

type Config struct {
	NATSURL string
}

// Resolve gets the path to message queue from the env variable and client name
func (c *Config) Resolve() (string, string, error) {
	if c.NATSURL != "" {
		return c.NATSURL, version.AppName, nil
	}

	err := envconfig.Process(version.AppName, c)
	if err != nil {
		panic(err)
	}

	if c.NATSURL == "" {
		return "", version.AppName, errors.New("cannot resolve nats url")
	}

	return c.NATSURL, version.AppName, nil
}
