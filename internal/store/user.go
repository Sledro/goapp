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
func (u *User) Create(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.NamedExec(createUserQuery, &u)
	tx.Commit()
	return nil
}

// Get - Gets a user
func (u *User) Get(db *sqlx.DB) (User, error) {
	user := User{}
	err := db.Get(&user, getUserQuery, u.ID, u.Username, u.Email)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

// Update - Updates a user
func (u *User) Update(userNew User, db *sqlx.DB) (User, error) {
	return *u, nil
}

// Delete - Deletes a user
func (u *User) Delete(db *sqlx.DB) error {
	return nil
}

// UserList - Gets a list of users
func (u *User) List(db *sqlx.DB) ([]User, error) {
	var userList []User
	// If search params then search users and return matching list
	if (User{}) != *u {
		err := db.Select(&userList, getUserListQuery, u.ID, u.Firstname, u.Lastname, u.Username, u.Email)
		if err != nil {
			return userList, err
		}
		return userList, nil
	}
	// No search params, return all
	err := db.Select(&userList, getUserListAllQuery)
	if err != nil {
		return userList, err
	}
	return userList, nil
}
