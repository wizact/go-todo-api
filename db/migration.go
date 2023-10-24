package dbmigration

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	_ "github.com/mattn/go-sqlite3"
)

type DBMigration struct{}

func (dbm *DBMigration) Start() error {
	fmt.Println("starting database migration")
	dir := "../../db"
	dbPath := filepath.Join(dir, "todo.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()

	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})

	d, err := bindata.WithInstance(s)

	if err != nil {
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("go-bindata", d, "sqlite3", driver)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil {
		return err
	}

	return nil
}
