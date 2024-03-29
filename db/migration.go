package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	_ "github.com/mattn/go-sqlite3"
)

type DBMigration struct{}

// Start initiates the migration process and close the db connection at the end of the process.
func (dbm *DBMigration) Start(dbPath string) error {
	fmt.Println("starting database migration")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()

	sd, err := getSource()

	if err != nil {
		return err
	}

	dd, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		return err
	}

	err = migrateSourceToDestination(sd, dd)

	if err != nil {
		return err
	}

	return nil
}

// getSource gets the source scripts from the autogenerated sourcefile.go.
func getSource() (source.Driver, error) {
	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})

	return bindata.WithInstance(s)
}

// migrateSourceToDestination migrates by executing no-applied source scripts to the destination instance.
func migrateSourceToDestination(sourceInstance source.Driver, databaseInstance database.Driver) error {
	m, err := migrate.NewWithInstance("go-bindata", sourceInstance, "sqlite3", databaseInstance)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		return err
	}

	return nil
}
