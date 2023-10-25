package main

import (
	"log"
	"path/filepath"

	migration "github.com/wizact/go-todo-api/db"
)

func main() {
	log.Println("database migration")
	dbm := migration.DBMigration{}

	dir := "../../db"
	dbPath := filepath.Join(dir, "todo.db")

	if err := dbm.Start(dbPath); err != nil {
		log.Fatalln(err)
	}
}
