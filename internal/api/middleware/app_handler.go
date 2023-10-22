package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

type AppHandler func(http.ResponseWriter, *http.Request) *hsm.AppError

// ServeHTTP to serve requests but respond with a friendly error message if any
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.

		log.Println(e.Error())

		w.WriteHeader(e.Code)
		ee := json.NewEncoder(w).Encode(&hsm.FriendlyError{Message: e.Message})
		if ee != nil {
			log.Fatal(ee.Error())
		}
	}
}
