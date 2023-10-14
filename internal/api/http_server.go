package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hndl "github.com/wizact/go-todo-api/internal/api/handlers"
)

const (
	healthCheckRoute = "/__health-check"
)

// StartServer starts the http server
func StartServer(address, port string, tls bool) {
	serverAddress := fmt.Sprintf("%s:%s", address, port)

	fmt.Println("Listening to requests from: " + serverAddress)

	router := mux.NewRouter()
	// router.Use(commonMiddleware)

	// HealthCheck route setup
	hcr := hndl.HealthCheckRoute{}
	hcr.SetupRoutes(healthCheckRoute, router)

	if tls {
		log.Fatal(http.ListenAndServeTLS(serverAddress,
			"certs/server.crt",
			"certs/server.key",
			router))
	} else {
		log.Fatal(http.ListenAndServe(serverAddress, router))
	}
}
