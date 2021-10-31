package store

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthStore struct {
	DB *sqlx.DB
}

type AuthStoreInterface interface {
	Login(user User) (User, error)
}

var AuthStoreInstance AuthStoreInterface = &AuthStore{}

var LoginQuery = `
SELECT * 
FROM users 
WHERE email=$1`

// Login - Returns user of provided email
func (a *AuthStore) Login(user User) (User, error) {
	u := User{}
	err := a.DB.Get(&u, LoginQuery, user.Email)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return u, errors.New("user does not exist")
		}
		return u, err
	}
	return u, nil
}
