package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sledro/goapp/api"
)

// Credentials - Holds login credentials
type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// handleAuthLogin - This function checks the users email and password
// match. If they matcn an oauth token is returned
func (s *server) handleAuthLogin(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}
	// Check the password and get the token
	token, err := s.services.AuthService.Login(credentials.Email, credentials.Password)
	if err != nil {
		api.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	api.JSON(w, http.StatusOK, map[string]interface{}{"token": token})
}
