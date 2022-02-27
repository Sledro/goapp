package services

import (
	"errors"

	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/auth"
)

type UserService struct {
	UserStore store.UserStoreInterface
}

type UserServiceInterface interface {
	Create(user store.User) (store.User, error)
	Get(user store.User) (store.User, error)
	Update(user store.User) (store.User, error)
	Delete(userID int) error
	List(user store.User) ([]store.User, error)
}

var UserServiceInstance UserServiceInterface = &UserService{}

// Create a user
func (s *UserService) Create(user store.User) (store.User, error) {
	// Check if user already exists
	u, _ := s.UserStore.Get(user)
	if u.ID > 0 {
		return store.User{}, errors.New("this user already exists")
	}

	// Hash password before we store it
	pass, err := auth.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = string(pass)

	// Store user
	user, err = s.UserStore.Create(user)
	if err != nil {
		return user, err
	}
	// Remove passowrd before returning
	user.Password = ""
	return user, nil
}

// Get a user
func (s *UserService) Get(user store.User) (store.User, error) {
	user, err := s.UserStore.Get(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Updates a user
func (s *UserService) Update(user store.User) (store.User, error) {
	user, err := s.UserStore.Update(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Deletes a user with given id
func (s *UserService) Delete(userID int) error {
	err := s.UserStore.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

// Get list of all users
func (s *UserService) List(user store.User) ([]store.User, error) {
	userList, err := s.UserStore.List(user)
	if err != nil {
		return []store.User{}, err
	}
	return userList, nil
}
