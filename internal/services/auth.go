package services

import (
	"github.com/sledro/golang-framework/internal/storage"
	"github.com/sledro/golang-framework/pkg/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login - Here we cheecek the user email exists and the passwords match.
// Finally ann auth token is created and returned
func Login(email, password, apiToken string, db *gorm.DB) (string, error) {
	// Check user with that email exists
	user, err := UserGet(storage.User{Email: email}, db)
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
