package store

import (
	"database/sql"
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
	Create(user User) (User, error)
	Get(user User) (User, error)
	Update(user User) (User, error)
	Delete(userID int) error
	List(user User) ([]User, error)
}

var UserStoreInstance UserStoreInterface = &UserStore{}

var CreateUserQuery = `
INSERT INTO users (firstname, lastname, username, password, email)
VALUES (?, ?, ?, ?, ?)`

var GetUserQuery = `
SELECT id, firstname, lastname, username, email
FROM users
WHERE id=? OR username=? OR email=?`

var updateUserQuery = `
UPDATE users 
SET firstname=?, lastname=?, username=?, email=?
WHERE id=?`

var deleteUserQuery = `
DELETE 
FROM users 
WHERE id=?`

var getUserListQuery = `
SELECT id, firstname, lastname, username, email
FROM users
WHERE id=? OR firstname=? OR lastname=? OR username=? OR email=?`

var getUserListAllQuery = `
SELECT id, firstname, lastname, username, email
FROM users`

// Create - Creates a user
func (s *UserStore) Create(user User) (User, error) {
	tx := s.DB.MustBegin()
	_, err := tx.Exec(CreateUserQuery, &user.Firstname, &user.Lastname, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}
	err = tx.Commit()
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get - Gets a user
func (s *UserStore) Get(user User) (User, error) {
	u := User{}
	err := s.DB.Get(&u, GetUserQuery, user.ID, user.Username, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}

// Update - Updates a user
func (s *UserStore) Update(user User) (User, error) {
	tx := s.DB.MustBegin()
	_, err := s.DB.Exec(updateUserQuery, user.Firstname, user.Lastname, user.Username, user.Email, user.ID)
	if err != nil {
		return user, err
	}
	err = tx.Commit()
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete - Deletes a user with given id
func (s *UserStore) Delete(userID int) error {
	tx := s.DB.MustBegin()
	_, err := s.DB.Exec(deleteUserQuery, userID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
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
