package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ErrFailedToConnectToDb = errors.New("failed to connect to database")
)

type SqliteConnection struct{}

func (slc SqliteConnection) Connection(cnf gorm.Config) (*gorm.DB, error) {

	dc := DbConfig{}
	dp, err := dc.GetDbPath()

	if err != nil {
		return nil, ErrFailedToConnectToDb
	}

	db, err := gorm.Open(sqlite.Open(dp), &cnf)

	if err != nil {
		return nil, ErrFailedToConnectToDb
	}

	return db, nil

}
