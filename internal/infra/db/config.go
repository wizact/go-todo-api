package db

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/internal/infra"
)

type DbConfig struct {
	DbPath string
}

// GetDbPath gets the path to sqlite database from the env variable
func (d *DbConfig) GetDbPath() (string, error) {
	if d.DbPath != "" {
		return d.DbPath, nil
	}

	err := envconfig.Process(infra.APPNAME, d)
	if err != nil {
		panic(err)
	}

	if d.DbPath == "" {
		return "", errors.New("cannot resolve database path")
	}

	return d.DbPath, nil
}
