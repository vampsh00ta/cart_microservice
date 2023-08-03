package coding

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func GetJwtHTTP(r *http.Request) (string, error) {

	authToken := r.Header.Get("Authorization")

	tokenSplited := strings.Split(authToken, " ")
	if len(tokenSplited) <= 1 {
		return "", errors.New("CODE_INVALID_AUTH_TOKEN")
	}
	rawToken := tokenSplited[1]

	return rawToken, nil
}

func DecodeJwt(secretToken string, user jwt.Claims, rawToken string) error {

	token, err := jwt.ParseWithClaims(rawToken, user, func(token *jwt.Token) (interface{}, error) {

		return []byte(secretToken), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("CODE_INVALID_AUTH_TOKEN")
	}
	return nil
}
