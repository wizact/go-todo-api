package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wizact/go-todo-api/internal/api/middleware"
	userModule "github.com/wizact/go-todo-api/internal/user"
	"github.com/wizact/go-todo-api/internal/user/interfaces/controller"
	httpModel "github.com/wizact/go-todo-api/internal/user/interfaces/model"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type UserRouteFactory struct {
}

func (urf UserRouteFactory) CreateUserRoute() UserRoute {

	return NewUserRoute(
		controller.NewUserController(
			userModule.NewUserModule(true).ResolveUserAccountUseCase()),
	)

}

type UserRoute struct {
	UserController controller.UserController
}

func NewUserRoute(userController controller.UserController) UserRoute {
	return UserRoute{
		UserController: userController,
	}
}

func (ur UserRoute) SetupRoutes(routePath string, router *mux.Router) {
	router.Handle(routePath, middleware.AppHandler(ur.RegisterUser())).Methods("POST")
}

// RegisterUser registers a user
func (ur UserRoute) RegisterUser() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {

		var u httpModel.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			return &hsm.AppError{ErrorObject: err, Message: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err = ur.UserController.RegisterNewUser(u)

		if e, a := err.(*hsm.AppError); !a {
			return e
		}

		json.NewEncoder(w).Encode(u)
		return nil
	}

	return fn
}
