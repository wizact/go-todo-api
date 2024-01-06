package main

import (
	"context"
	"fmt"

	api "github.com/wizact/go-todo-api/internal/api"
	"github.com/wizact/go-todo-api/internal/infra"
	"github.com/wizact/yacli"
)

func main() {
	fmt.Printf("GO DDD And Clean Architecture Example")
	fmt.Println()

	app := yacli.NewApplication()

	app.Name = infra.APPNAME
	app.Description = "Go DDD & Clean Architecture API Example"

	app.AddCommand(&api.StartServerCommand{})

	ctx := context.Background()

	app.Run(ctx)
}
