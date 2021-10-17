package services

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/auth"
)

// UserCreate - Create a user
func UserCreate(user store.User, db *sqlx.DB) error {
	pass, err := auth.HashPassword(user.Password)
	user.Password = string(pass)
	if err != nil {
		return nil
	}
	// Check if user already exists
	_, err = user.Get(db)
	if err == nil {
		return errors.New("username or email already exists")
	}

	err = user.Create(db)
	if err != nil {
		return nil
	}
	return nil
}

// UserCreate - Get a user
func UserGet(user store.User, db *sqlx.DB) (store.User, error) {
	user, err := user.Get(db)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// UserUpdate - Updates a user
func UserUpdate(user store.User, userNew store.User, db *sqlx.DB) (store.User, error) {
	user, err := user.Update(userNew, db)
	if err != nil {
		return store.User{}, err
	}
	return user, nil
}

// UserDelete - Deletes a user
func UserDelete(user store.User, db *sqlx.DB) error {
	err := user.Delete(db)
	if err != nil {
		return err
	}
	return nil
}

// UserCreate - Get list of all users
func UserList(user store.User, db *sqlx.DB) ([]store.User, error) {
	userList, err := user.List(db)
	if err != nil {
		return []store.User{}, err
	}
	return userList, nil
}
