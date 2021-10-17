package store

import (
	"github.com/jmoiron/sqlx"
)

type AuthStore struct {
	DB *sqlx.DB
}

type AuthStoreInterface interface {
	Login(user User) (User, error)
}

var AuthStoreInstance AuthStoreInterface = &AuthStore{}

var loginQuery = `
SELECT * 
FROM users 
WHERE email=$1`

// Login - Returns user of provided email
func (a *AuthStore) Login(user User) (User, error) {
	u := User{}
	err := a.DB.Get(&u, loginQuery, user.Email)
	if err != nil {
		return u, err
	}
	return u, nil
}
