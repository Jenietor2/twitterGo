package jwt

import (
	"curso_go/twitterGo/models"
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUser string

func ProccessToken(token string, JWTSing string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSing)

	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token inválido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		//TODO
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token invalid")
	}

	return &claims, false, string(""), err
}
