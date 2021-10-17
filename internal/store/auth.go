package store

import (
	"github.com/jmoiron/sqlx"
)

var loginQuery = `
SELECT * 
FROM users 
WHERE email=$1`

// Login - Returns user of provided email
func Login(user User, db *sqlx.DB) error {
	u := User{}
	err := db.Get(&u, loginQuery, user.Email)
	if err != nil {
		return err
	}
	return nil
}
