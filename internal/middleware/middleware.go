package middleware

import (
	"errors"
	"net/http"

	"github.com/sledro/golang-framework/api"
	"github.com/sledro/golang-framework/pkg/auth"
)

// Headers - Sets HTTP headers such as content type
// and CORS configurations
func Headers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

// Auth - Sets Auth checek on a route. Checks JWT
// token is valid or returns status 404 Unauthorized
func Auth(next http.HandlerFunc, apiToken string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r, apiToken)
		if err != nil {
			api.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
