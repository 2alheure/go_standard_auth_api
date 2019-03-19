package helpers

import (
	"net/http"
	"strings"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenContent struct {
	UserID				int
	jwt.StandardClaims
}

func CheckToken(r *http.Request) (*TokenContent, *StdErr) {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return nil, &StdErr{"Missing auth token.", 400}
	}

	splitted := strings.Split(tokenHeader, " ")	// Token normally comes in form "Bearer <token>"
	if len(splitted) != 2 {
		return nil, &StdErr{"Malformed auth token.", 400}
	}

	tokenContent := &TokenContent{}
	token, err := jwt.ParseWithClaims(splitted[1], tokenContent, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if !token.Valid || err != nil {
		return nil, &StdErr{"Invalid token.", 409}
	}

	return tokenContent, nil
}

func CreateToken(userID int) (string, error) {
	tk := &TokenContent{UserID: userID}
	tokenSign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	token, err := tokenSign.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	return "Bearer "+token, err
}
