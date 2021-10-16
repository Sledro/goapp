package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"errors"

	validator "github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sledro/golang-framework/api"
	"github.com/sledro/golang-framework/internal/services"
	"github.com/sledro/golang-framework/internal/storage"
)

// handleUserCreate - Create a user
func (s *server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Create user object
	user := storage.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Validate data
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Create user
	err = services.UserCreate(user, s.db)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Sort the data to be returned
	sortData := map[string]interface{}{
		"status": "success",
	}
	api.JSON(w, http.StatusCreated, sortData)
}

// handleUserView - Get a user
func (s *server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Create user object
	user := storage.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Remove password from user so we dont seach on that field
	if user.Password != "" {
		api.ERROR(w, http.StatusUnprocessableEntity, errors.New("can not use password field here"))
		return
	}

	// Get user
	user, err = services.UserGet(user, s.db)
	user.Password = ""
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, user)
}

// handleUserList - Updates a user
func (s *server) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userIDString := mux.Vars(r)["id"]
	userIDInt, _ := strconv.Atoi(userIDString)
	user := storage.User{ID: userIDInt}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Create user object
	userNew := storage.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		api.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Update user
	userUpdated, err := services.UserUpdate(user, userNew, s.db)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, userUpdated)
}

// handleUserList - Deletes a user
func (s *server) handleUserDelete(w http.ResponseWriter, r *http.Request) {
	// Parse path var
	userIDString := mux.Vars(r)["id"]
	userIDInt, _ := strconv.Atoi(userIDString)
	user := storage.User{ID: userIDInt}

	// Delete user
	err := services.UserDelete(user, s.db)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusOK, "success")
}

// handleUserList - Get a list of all users
func (s *server) handleUserList(w http.ResponseWriter, r *http.Request) {
	// Create user
	userList, err := services.UserList(s.db)
	if err != nil {
		api.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	api.JSON(w, http.StatusCreated, userList)
}