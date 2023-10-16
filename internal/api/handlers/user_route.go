package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mw "github.com/wizact/go-todo-api/internal/api/middleware"
	um "github.com/wizact/go-todo-api/internal/user"
	uc "github.com/wizact/go-todo-api/internal/user/interfaces/controller"
	hm "github.com/wizact/go-todo-api/pkg/http-models"
)

type UserRouteFactory struct {
}

func (urf UserRouteFactory) CreateUserRoute() UserRoute {

	return NewUserRoute(
		uc.NewUserController(
			um.NewUserModule(true, true).ResolveUserAccountUseCase()),
	)

}

type UserRoute struct {
	UserController uc.UserController
}

func NewUserRoute(userController uc.UserController) UserRoute {
	return UserRoute{
		UserController: userController,
	}
}

func (ur UserRoute) SetupRoutes(routePath string, router *mux.Router) {
	router.Handle(routePath, mw.AppHandler(ur.RegisterUser())).Methods("GET")
}

// RegisterUser registers a user
func (ur UserRoute) RegisterUser() mw.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hm.AppError {
		ur.UserController.RegisterNewUser()
		json.NewEncoder(w).Encode("OK")
		return nil
	}

	return fn
}
