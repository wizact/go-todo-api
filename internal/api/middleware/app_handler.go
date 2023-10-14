package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	hm "github.com/wizact/go-todo-api/pkg/http-models"
)

type AppHandler func(http.ResponseWriter, *http.Request) *hm.AppError

// ServeHTTP to serve requests but respond with a friendly error message if any
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.

		log.Println(e.Error)

		w.WriteHeader(e.Code)
		ee := json.NewEncoder(w).Encode(&hm.FriendlyError{Message: e.Message})
		if ee != nil {
			log.Fatal(ee.Error())
		}
	}
}
