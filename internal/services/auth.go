package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/sledro/goapp/internal/store"
	"github.com/sledro/goapp/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DB          *sqlx.DB
	JWTSecret   string
	UserService UserServiceInterface
	AuthStore   store.AuthStoreInterface
}

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
}

var AuthServiceInstance AuthServiceInterface = &AuthService{}

// Here we checek the user email exists and the passwords match.
// Finally an auth token is created and returned
func (s *AuthService) Login(email, password string) (string, error) {
	// Check user with that email exists
	user, err := s.AuthStore.Login(store.User{Email: email})
	if err != nil {
		return "", err
	}

	// Verify the password is valid
	err = auth.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// Get auth token
	token, err := auth.CreateToken(user.ID, s.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
