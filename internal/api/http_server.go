package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hndl "github.com/wizact/go-todo-api/internal/api/handlers"
	comms "github.com/wizact/go-todo-api/pkg/communication"
)

const (
	HealthCheckRoute = "/__health-check"
	UserRoute        = "/users"
)

// StartServer starts the http server
func StartServer(address, port string, tls bool) {
	serverAddress := fmt.Sprintf("%s:%s", address, port)

	fmt.Println("Listening to requests from: " + serverAddress)

	router := mux.NewRouter()
	// router.Use(commonMiddleware)

	// Register services
	registerBackgroundServices()

	// Register all the routes
	registerRoutes(router)

	if tls {
		log.Fatal(http.ListenAndServeTLS(serverAddress,
			"certs/server.crt",
			"certs/server.key",
			router))
	} else {
		log.Fatal(http.ListenAndServe(serverAddress, router))
	}
}

func registerBackgroundServices() {
	comms.NewCommsModule(false)
}

func registerRoutes(router *mux.Router) {
	// HealthCheck route setup
	hcr := hndl.HealthCheckRoute{}
	hcr.SetupRoutes(HealthCheckRoute, router)

	// User route setup
	ur := hndl.UserRouteFactory{}.CreateUserRoute()
	ur.SetupRoutes(UserRoute, router)

}
