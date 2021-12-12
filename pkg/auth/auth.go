package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type customClaims struct {
	UserID int  `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.StandardClaims
}

// CreateToken - Creates abd returns a new JTW token for the
// given userID. Token expires after 1 hour
func CreateToken(userID int, apiSecret string) (string, error) {
	// Create the Claims
	claims := customClaims{
		userID,
		true,
		jwt.StandardClaims{
			Id:        strconv.FormatInt(time.Now().Unix()+int64(userID), 10),
			Issuer:    "http://localhost.com:8080",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(apiSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil

}

// TokenValid - Checks if the token passed in the request is
// valid. If not valid an error will reee returned
func TokenValid(r *http.Request, apiSecret string) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invliad token")
	}
	return nil
}

// extractToken - Gets a token from the req header
func extractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// HashPassword - Returns the bcrypt hash of the password
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// HashPassword - Verifies the bcrypt hash of the password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
