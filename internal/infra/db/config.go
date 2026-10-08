package db

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/pkg/version"
)

type Config struct {
	DBPath string
}

// ResolvePath gets the path to sqlite database from the env variable
func (c *Config) ResolvePath() (string, error) {
	if c.DBPath != "" {
		return c.DBPath, nil
	}

	err := envconfig.Process(version.AppName, c)
	if err != nil {
		panic(err)
	}

	if c.DBPath == "" {
		return "", errors.New("cannot resolve database path")
	}

	return c.DBPath, nil
}
