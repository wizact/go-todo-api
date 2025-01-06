package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type AppHandler func(http.ResponseWriter, *http.Request) *hsm.AppError

func (fm AppHandler) Config(requireAuth bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if requireAuth {
			if authz := fm.authorise(w, r); !authz {
				// terminate the request if the user is not authorized
				return
			}
		}

		fm.serveHTTP(w, r)
	}
}

// ServeHTTP to serve requests but respond with a friendly error message if any
func (fn AppHandler) serveHTTP(w http.ResponseWriter, r *http.Request) {

	if e := fn(w, r); e != nil { // e is *appError, not os.Error.

		log.Println(e.Error())

		w.WriteHeader(e.Code)
		ee := json.NewEncoder(w).Encode(&hsm.FriendlyError{Message: e.SanitisedMessage})
		if ee != nil {
			log.Fatal(ee.Error())
		}
	}
}

// type AuthHandler func(http.ResponseWriter, *http.Request) *hsm.AppError

// authorise to check if the user is authorized to access the resource
func (fn AppHandler) authorise(w http.ResponseWriter, r *http.Request) bool {
	if au := r.Header.Get("Authorization"); au != "" && au == "'Bearer bearer token'" {
		// TODO: Implement a proper authorization mechanism
		return true
	}

	log.Println("Unauthorized request")
	w.WriteHeader(http.StatusUnauthorized)
	ee := json.NewEncoder(w).Encode(&hsm.FriendlyError{Message: "Unauthorized request"})
	if ee != nil {
		log.Fatal(ee.Error())
	}
	return false
}
