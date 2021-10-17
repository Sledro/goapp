package store

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// User - An app user
type User struct {
	ID        int    `json:"id" db:"id"`
	Firstname string `json:"firstname" db:"firstname" validate:"required,min=3,max=100"`
	Lastname  string `json:"lastname" db:"lastname" validate:"required,min=3,max=100"`
	Username  string `json:"username" db:"username" validate:"required,min=3,max=30"`
	Password  string `json:"password,omitempty" db:"password" validate:"required,min=8,max=30"`
	Email     string `json:"email" db:"email" validate:"required,email"`
}

type UserStore struct {
	DB *sqlx.DB
}

type UserStoreInterface interface {
	Create(user User) error
	Get(user User) (User, error)
	Update(userNew User) (User, error)
	Delete(user User) error
	List(user User) ([]User, error)
}

var UserStoreInstance UserStoreInterface = &UserStore{}

var createUserQuery = `
INSERT INTO users (firstname, lastname, username, password, email)
VALUES (:firstname, :lastname, :username, :password, :email)`

var getUserQuery = `
SELECT id, firstname, lastname, username, email
FROM users
WHERE id=$1 OR username=$2 OR email=$3`

var getUserListQuery = `
SELECT id, firstname, lastname, username, email
FROM users
WHERE id=$1 OR firstname=$2 OR lastname=$3 OR username=$4 OR email=$5`

var getUserListAllQuery = `
SELECT id, firstname, lastname, username, email
FROM users`

// Create - Creates a user
func (s *UserStore) Create(user User) error {
	tx := s.DB.MustBegin()
	tx.NamedExec(createUserQuery, &user)
	tx.Commit()
	return nil
}

// Get - Gets a user
func (s *UserStore) Get(user User) (User, error) {
	userRes := User{}
	err := s.DB.Get(&userRes, getUserQuery, user.ID, user.Username, user.Email)
	if err != nil {
		return userRes, errors.New("user not found")
	}
	return userRes, nil
}

// Update - Updates a user
func (s *UserStore) Update(userNew User) (User, error) {
	return userNew, nil
}

// Delete - Deletes a user
func (s *UserStore) Delete(user User) error {
	return nil
}

// UserList - Gets a list of users
func (s *UserStore) List(user User) ([]User, error) {
	var userList []User
	// If search params then search users and return matching list
	if (User{}) != user {
		err := s.DB.Select(&userList, getUserListQuery, user.ID, user.Firstname, user.Lastname, user.Username, user.Email)
		if err != nil {
			return userList, err
		}
		return userList, nil
	}
	// No search params, return all
	err := s.DB.Select(&userList, getUserListAllQuery)
	if err != nil {
		return userList, err
	}
	return userList, nil
}
