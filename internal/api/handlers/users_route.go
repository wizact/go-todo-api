package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/wizact/go-todo-api/internal/api/middleware"
	userModule "github.com/wizact/go-todo-api/internal/user"
	controller "github.com/wizact/go-todo-api/internal/user/adapters/controllers"
	httpModel "github.com/wizact/go-todo-api/internal/user/adapters/controllers/models"
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
	router.Handle(routePath+"/{id}", middleware.AppHandler(ur.GetUserById())).Methods("GET")
}

// RegisterUser registers a user
func (ur UserRoute) RegisterUser() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {

		var u httpModel.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			return &hsm.AppError{ErrorObject: err, Message: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err = ur.UserController.RegisterNewUser(context.TODO(), u)

		e, a := err.(*hsm.AppError)
		if e != nil && a {
			return e
		}

		w.Header().Add("location", fmt.Sprintf("/users/%v", u.UserID))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
		return nil
	}

	return fn
}

// GetUserById retrieves a user by it's Id
func (ur UserRoute) GetUserById() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {

		var uid uuid.UUID
		var err error
		if uid, err = uuid.Parse(mux.Vars(r)["id"]); err != nil {
			return &hsm.AppError{ErrorObject: err, Message: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err := ur.UserController.GetUserById(context.TODO(), uid)
		e, a := err.(*hsm.AppError)
		if e != nil && a {
			return e
		}

		json.NewEncoder(w).Encode(u)
		return nil
	}

	return fn
}
