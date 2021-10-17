package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

// Login - Here we checek the user email exists and the passwords match.
// Finally an auth token is created and returned
func Login(email, password, apiToken string, db *sqlx.DB) (string, error) {
	// Check user with that email exists
	user, err := UserGet(store.User{Email: email}, db)
	if err != nil {
		return "", err
	}

	// Verify the password is valid
	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Get auth token
	token, err := auth.CreateToken(user.ID, apiToken)
	if err != nil {
		return "", err
	}

	return token, nil
}
