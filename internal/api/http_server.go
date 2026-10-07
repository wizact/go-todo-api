package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	hndl "github.com/wizact/go-todo-api/internal/api/handlers"
	usermodule "github.com/wizact/go-todo-api/internal/user"
	applicationport "github.com/wizact/go-todo-api/internal/user/ports/applications"
	usecaseport "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
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
	userModule := usermodule.NewUserModule(true)
	userAccount := userModule.UserAccountUseCase()
	registration := userModule.UserRegistrationAppService()

	// Register services
	registerBackgroundServices(registration)

	// Register all the routes
	registerRoutes(router, userAccount, registration)

	if tls {
		log.Fatal(http.ListenAndServeTLS(serverAddress,
			"certs/server.crt",
			"certs/server.key",
			router))
	} else {
		log.Fatal(http.ListenAndServe(serverAddress, router))
	}
}

func registerBackgroundServices(registration applicationport.Registration) {
	if _, err := comms.NewCommsModule(comms.Config{}, registration); err != nil {
		panic(err)
	}
}

func registerRoutes(
	router *mux.Router,
	userAccount usecaseport.UserAccountUseCase,
	registration applicationport.Registration,
) {
	// HealthCheck route setup
	hcr := hndl.HealthCheckRoute{}
	hcr.SetupRoutes(HealthCheckRoute, router)

	// User route setup
	ur := hndl.UserRouteFactory{
		UserAccountUseCase: userAccount,
		Registration:       registration,
	}.CreateUserRoute()
	ur.SetupRoutes(UserRoute, router)

}
