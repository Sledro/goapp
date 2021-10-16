package services

import (
	"github.com/sledro/golang-framework/internal/storage"
	"github.com/sledro/golang-framework/pkg/auth"
	"gorm.io/gorm"
)

// UserCreate - Create a user
func UserCreate(user storage.User, db *gorm.DB) error {
	pass, err := auth.HashPassword(user.Password)
	user.Password = string(pass)
	if err != nil {
		return nil
	}
	err = storage.UserCreate(user, db)
	if err != nil {
		return nil
	}
	return nil
}

// UserCreate - Get a user
func UserGet(user storage.User, db *gorm.DB) (storage.User, error) {
	user, err := storage.UserGet(user, db)
	if err != nil {
		return storage.User{}, err
	}
	return user, nil
}

// UserUpdate - Updates a user
func UserUpdate(user storage.User, userNew storage.User, db *gorm.DB) (storage.User, error) {
	user, err := storage.UserUpdate(user, userNew, db)
	if err != nil {
		return storage.User{}, err
	}
	return user, nil
}

// UserDelete - Deletes a user
func UserDelete(user storage.User, db *gorm.DB) error {
	err := storage.UserDelete(user, db)
	if err != nil {
		return err
	}
	return nil
}

// UserCreate - Get list of all users
func UserList(db *gorm.DB) ([]storage.User, error) {
	userList, err := storage.UserList(db)
	if err != nil {
		return []storage.User{}, err
	}
	return userList, nil
}
