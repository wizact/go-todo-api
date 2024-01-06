package main

import (
	"log"

	migration "github.com/wizact/go-todo-api/db"
	"github.com/wizact/go-todo-api/internal/infra/config"
)

func main() {
	log.Println("database migration")
	dbm := migration.DBMigration{}

	dc := &config.DbConfig{}

	dp, err := dc.GetDbPath()

	if err != nil {
		log.Fatalf("error resolving db path: %v", err.Error())
	}

	if err = dbm.Start(dp); err != nil {
		log.Fatalln(err)
	}
}
