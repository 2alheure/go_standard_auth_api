package helpers

import (
	"net/http"
	"strings"
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenContent struct {
	UserID				int
	jwt.StandardClaims
}

func CheckToken(r *http.Request) (*TokenContent, error) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return nil, errors.New("Missing auth token.")
	}

	splitted := strings.Split(tokenHeader, " ")	// Token normally comes in form "Bearer <token>"
	if len(splitted) != 2 {
		return nil, errors.New("Malformed auth token.")
	}

	tokenContent := &TokenContent{}
	token, err := jwt.ParseWithClaims(splitted[1], tokenContent, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		return nil, errors.New("Error while parsing token.")
	}

	if !token.Valid {
		return nil, errors.New("Invalid token.")
	}

	return tokenContent, nil
}

