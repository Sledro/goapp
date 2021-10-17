package store

import (
	"github.com/jmoiron/sqlx"
)

// User - An app user
type User struct {
	ID        int    `json:"user_id"`
	Firstname string `json:"firstname" validate:"required,min=3,max=100"`
	Lastname  string `json:"lastname" validate:"required,min=3,max=100"`
	Username  string `json:"username" validate:"required,min=3,max=30"`
	Password  string `json:"password,omitempty" validate:"required,min=8,max=30"`
	Email     string `json:"email" validate:"required,email"`
}

// HomeAddress - Users home address

var createUserQuery = `
INSERT INTO users (firstname, lastname, username, password, email) 
VALUES (:firstname, :lastname, :username, :password, :email)`

var getUserQuery = `
SELECT (firstname, lastname, username, email)
FROM users
WHERE id=$1`

var getUserListQuery = `
SELECT (firstname, lastname, username, email)
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
	err := db.Get(&user, getUserQuery, u.ID)
	if err != nil {
		return user, err
	}
	return *u, nil
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
func UserList(db *sqlx.DB) ([]User, error) {
	var userList []User
	err := db.Select(&userList, getUserListQuery)
	if err != nil {
		return userList, err
	}
	return userList, nil
}
