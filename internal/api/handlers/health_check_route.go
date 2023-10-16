package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mw "github.com/wizact/go-todo-api/internal/api/middleware"
	hm "github.com/wizact/go-todo-api/pkg/http-models"
)

type HealthCheckRoute struct {
}

func NewHealthCheckRoute() HealthCheckRoute {
	return HealthCheckRoute{}
}

func (hcr HealthCheckRoute) SetupRoutes(routePath string, router *mux.Router) {
	router.Handle(routePath, mw.AppHandler(hcr.GetHealthCheck())).Methods("GET")
}

// GetHealthCheck returns OK when is called
func (hcr HealthCheckRoute) GetHealthCheck() mw.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hm.AppError {
		json.NewEncoder(w).Encode("OK")
		return nil
	}

	return fn
}
