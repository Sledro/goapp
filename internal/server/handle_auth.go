package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sledro/golang-framework/api"
	"github.com/sledro/golang-framework/internal/services"
)

// Credentials - Holds login credentials
type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// handleAuthLogin - This function checks the users email and password
// match. If they matcn an oauth token is returned
func (s *server) handleAuthLogin(w http.ResponseWriter, r *http.Request) {
	// Read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Parse credentials
	credentials := Credentials{}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Check the password and get the token
	token, err := services.Login(credentials.Email, credentials.Password, s.secrets.JWTSecret, s.db)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// Sort the data to be returned
	sortData := map[string]string{
		"token": token,
	}

	api.JSON(w, http.StatusOK, sortData)
}
