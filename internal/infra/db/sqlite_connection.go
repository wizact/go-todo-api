package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrFailedToResolveDbConnStr = errors.New("failed to resolve db connection string")
	ErrFailedToConnectToDb      = errors.New("failed to connect to database")
)

type SqliteConnection struct {
	connectionString string
}

// NewSqliteConnection create a new sqlite connection but it does not connect to it.
// If connectionString is not provided, then it resolves it from env variables.
func NewSqliteConnection(connectionString string) (*SqliteConnection, error) {

	dp, err := resolveConnectionString(connectionString)

	if err != nil {
		return nil, err
	}

	return &SqliteConnection{connectionString: dp}, nil
}

func resolveConnectionString(connectionString string) (string, error) {
	if connectionString != "" {
		return connectionString, nil
	}

	dc := DbConfig{}
	dp, err := dc.GetDbPath()

	if err != nil {
		return "", ErrFailedToResolveDbConnStr
	}

	return dp, nil
}

func (slc *SqliteConnection) Open(cnf gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(slc.connectionString), &cnf)

	if err != nil {
		return nil, ErrFailedToConnectToDb
	}

	return db, nil
}
