package main

import (
	"log"

	migration "github.com/wizact/go-todo-api/db"
	infradb "github.com/wizact/go-todo-api/internal/infra/db"
)

func main() {
	log.Println("database migration")
	dbm := migration.DBMigration{}

	dp, err := (&infradb.DbConfig{}).GetDbPath()
	if err != nil {
		log.Fatalf("error resolving db path: %v", err.Error())
	}

	if err = dbm.Start(dp); err != nil {
		if err.Error() == "no change" {
			log.Println(err.Error())
			return
		}
		log.Fatalln(err)
	}
}
