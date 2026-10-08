package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrResolveConnectionString = errors.New("failed to resolve db connection string")
	ErrConnectDB               = errors.New("failed to connect to database")
)

type SQLiteConnection struct {
	connectionString string
}

// NewSQLiteConnection create a new sqlite connection but it does not connect to it.
// If connectionString is not provided, then it resolves it from env variables.
func NewSQLiteConnection(connectionString string) (*SQLiteConnection, error) {

	dp, err := resolveConnectionString(connectionString)

	if err != nil {
		return nil, err
	}

	return &SQLiteConnection{connectionString: dp}, nil
}

func resolveConnectionString(connectionString string) (string, error) {
	if connectionString != "" {
		return connectionString, nil
	}

	dc := Config{}
	dp, err := dc.ResolvePath()

	if err != nil {
		return "", ErrResolveConnectionString
	}

	return dp, nil
}

func (s *SQLiteConnection) Open(cnf gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.connectionString), &cnf)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrConnectDB, err)
	}

	return db, nil
}
