package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/wizact/go-todo-api/internal/api/middleware"
	controller "github.com/wizact/go-todo-api/internal/user/adapters/controller"
	httpModel "github.com/wizact/go-todo-api/internal/user/adapters/controller/model"
	usecaseport "github.com/wizact/go-todo-api/internal/user/ports/input/usecase"
	applicationport "github.com/wizact/go-todo-api/internal/user/ports/service"
	hsm "github.com/wizact/go-todo-api/pkg/httpservermodel"
)

type UserRouteFactory struct {
	UserAccountUseCase usecaseport.UserAccountUseCase
	Registration       applicationport.Registration
}

func (f UserRouteFactory) CreateUserRoute() UserRoute {
	return NewUserRoute(
		controller.New(
			f.UserAccountUseCase,
			f.Registration,
		),
	)
}

type UserRoute struct {
	Controller controller.Controller
}

func NewUserRoute(userController controller.Controller) UserRoute {
	return UserRoute{
		Controller: userController,
	}
}

func (ur UserRoute) SetupRoutes(routePath string, router *mux.Router) {
	router.Handle(routePath, middleware.AppHandler(ur.RegisterUser()).Config(false)).Methods("POST")
	router.Handle(routePath+"/verify-registration", middleware.AppHandler(ur.VerifyRegistration()).Config(false)).Methods("POST")

	router.Handle(routePath+"/{id}", middleware.AppHandler(ur.FetchUserByID()).Config(true)).Methods("GET")
}

func (ur UserRoute) VerifyRegistration() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {
		var uid uuid.UUID
		var token string
		var err error
		if uid, err = uuid.Parse(r.URL.Query().Get("uid")); err != nil || uid == uuid.Nil {
			return &hsm.AppError{ErrorObject: err, SanitisedMessage: "Bad Request", Code: http.StatusBadRequest}
		}

		if token = r.URL.Query().Get("token"); token == "" {
			return &hsm.AppError{ErrorObject: err, SanitisedMessage: "Bad Request", Code: http.StatusBadRequest}
		}

		err = ur.Controller.VerifyUserRegistration(r.Context(), uid, token)
		e, a := err.(*hsm.AppError)
		if e != nil && a {
			return e
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}

	return fn
}

// RegisterUser registers a user
func (ur UserRoute) RegisterUser() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {

		var u httpModel.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			return &hsm.AppError{ErrorObject: err, SanitisedMessage: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err = ur.Controller.RegisterNewUser(r.Context(), u)

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

// FetchUserByID retrieves a user by its ID
func (ur UserRoute) FetchUserByID() middleware.AppHandler {
	fn := func(w http.ResponseWriter, r *http.Request) *hsm.AppError {

		var uid uuid.UUID
		var err error
		if uid, err = uuid.Parse(mux.Vars(r)["id"]); err != nil {
			return &hsm.AppError{ErrorObject: err, SanitisedMessage: "Bad Request", Code: http.StatusBadRequest}
		}

		u, err := ur.Controller.FetchUserByID(r.Context(), uid)
		e, a := err.(*hsm.AppError)
		if e != nil && a {
			return e
		}

		json.NewEncoder(w).Encode(u)
		return nil
	}

	return fn
}
