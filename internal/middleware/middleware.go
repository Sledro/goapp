package middleware

import (
	"errors"
	"net/http"

	"github.com/sledro/goapp/api"
	"github.com/sledro/goapp/pkg/auth"
)

// Sets Auth checek on a route. Checks JWT token
// is valid or returns status 404 Unauthorized
func Auth(next http.HandlerFunc, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r, jwtSecret)
		if err != nil {
			api.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
