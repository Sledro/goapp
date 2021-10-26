package middleware

import (
	"errors"
	"net/http"

	"github.com/sledro/goapp/api"
	"github.com/sledro/goapp/pkg/auth"
)

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
