package main

import (
	"log"

	migration "github.com/wizact/go-todo-api/db"
)

func main() {
	log.Println("database migration")
	dbm := migration.DBMigration{}
	if err := dbm.Start(); err != nil {
		log.Fatalln(err)
	}
}
