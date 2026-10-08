package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mw "github.com/wizact/go-todo-api/internal/api/middleware"
	hsm "github.com/wizact/go-todo-api/pkg/httpservermodel"
)

type HealthCheckRoute struct {
}

func NewHealthCheckRoute() HealthCheckRoute {
	return HealthCheckRoute{}
}

func (h HealthCheckRoute) SetupRoutes(routePath string, router *mux.Router) {
	router.Handle(routePath, mw.AppHandler(h.HealthCheck()).Config(false)).Methods("GET")
}

// HealthCheck returns OK when is called
func (h HealthCheckRoute) HealthCheck() mw.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {
		json.NewEncoder(w).Encode("OK")
		return nil
	}

	return fn
}
