package api

import (
	"fmt"
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
func StartServer(address, port string, tls bool, config Config) error {
	serverAddress := fmt.Sprintf("%s:%s", address, port)

	fmt.Println("Listening to requests from: " + serverAddress)

	router := mux.NewRouter()
	// router.Use(commonMiddleware)
	userModule := usermodule.NewUserModule(true, config.registrationConfig())
	userAccount := userModule.UserAccountUseCase()
	registration := userModule.UserRegistrationAppService()

	// Register services
	if err := registerBackgroundServices(registration, config.communicationConfig()); err != nil {
		return fmt.Errorf("register background services: %w", err)
	}

	// Register all the routes
	registerRoutes(router, userAccount, registration)

	if tls {
		if err := http.ListenAndServeTLS(serverAddress,
			"certs/server.crt",
			"certs/server.key",
			router); err != nil {
			return fmt.Errorf("serve HTTPS: %w", err)
		}
		return nil
	}
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		return fmt.Errorf("serve HTTP: %w", err)
	}
	return nil
}

func registerBackgroundServices(registration applicationport.Registration, config comms.Config) error {
	if _, err := comms.NewCommsModule(config, registration); err != nil {
		return fmt.Errorf("start communication module: %w", err)
	}
	return nil
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
