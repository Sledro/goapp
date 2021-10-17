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
	Create(user store.User) error
	Get(user store.User) (store.User, error)
	Update(user store.User) (store.User, error)
	Delete(user store.User) error
	List(user store.User) ([]store.User, error)
}

var UserServiceInstance UserServiceInterface = &UserService{}

// Create - Create a user
func (s *UserService) Create(user store.User) error {
	pass, err := auth.HashPassword(user.Password)
	user.Password = string(pass)
	if err != nil {
		return nil
	}
	// Check if user already exists
	_, err = s.UserStore.Get(user)
	if err == nil {
		return errors.New("username or email already exists")
	}

	err = s.UserStore.Create(user)
	if err != nil {
		return nil
	}
	return nil
}

// Get - Get a user
func (s *UserService) Get(user store.User) (store.User, error) {
	user, err := s.UserStore.Get(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Update - Updates a user
func (s *UserService) Update(user store.User) (store.User, error) {
	user, err := s.UserStore.Update(user)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// Delete - Deletes a user
func (s *UserService) Delete(user store.User) error {
	err := s.UserStore.Delete(user)
	if err != nil {
		return err
	}
	return nil
}

// List - Get list of all users
func (s *UserService) List(user store.User) ([]store.User, error) {
	userList, err := s.UserStore.List(user)
	if err != nil {
		return []store.User{}, err
	}
	return userList, nil
}
