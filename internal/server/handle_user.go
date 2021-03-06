package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sledro/goapp/api"
	"github.com/sledro/goapp/internal/store"
)

// Creates a user
func (s *server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Create user
	user, err = s.services.UserService.Create(user)
	if err != nil {
		api.ERROR(w, http.StatusConflict, err)
		return
	}

	api.JSON(w, http.StatusCreated, user)
}

// Gets a user
func (s *server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	// Read the request
	var user store.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Get user
	u, err := s.services.UserService.Get(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, u)
}

// Update a user
func (s *server) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Read the request
	var user store.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("could not decode JSON body"))
		return
	}

	// Set userID that we want to update
	user.ID = userID

	// Update user
	user, err = s.services.UserService.Update(user)
	if err != nil {
		s.log.Error(err)
		api.ERROR(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	api.JSON(w, http.StatusOK, user)
}

// Deletes a user
func (s *server) handleUserDelete(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Delete user
	err = s.services.UserService.Delete(userID)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, map[string]interface{}{"success": true})
}

// Gets a list of all users
func (s *server) handleUserList(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Create user object
	user := store.User{}

	// If request had payload parse it
	if len(body) > 0 {
		err = json.Unmarshal(body, &user)
		if err != nil {
			api.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
	}

	// List users
	userList, err := s.services.UserService.List(user)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if len(userList) > 0 {
		api.JSON(w, http.StatusCreated, userList)
	} else {
		api.JSON(w, http.StatusOK, map[string]interface{}{"error": "no users"})
	}

}
