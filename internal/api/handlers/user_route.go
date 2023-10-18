package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	mw "github.com/wizact/go-todo-api/internal/api/middleware"
	um "github.com/wizact/go-todo-api/internal/user"
	uc "github.com/wizact/go-todo-api/internal/user/interfaces/controller"
	httpmodel "github.com/wizact/go-todo-api/internal/user/interfaces/model"
	hm "github.com/wizact/go-todo-api/pkg/http-model"
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

		var u httpmodel.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			return &hm.AppError{ErrorObject: err, Message: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err = ur.UserController.RegisterNewUser(u)

		if e, a := err.(*hm.AppError); a {
			return e
		}

		json.NewEncoder(w).Encode(u)
		return nil
	}

	return fn
}
