package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	Id         string
	Expiration time.Time
}

func (authClaims *AuthClaims) ApplyClaims(token *jwt.Token) {
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = authClaims.Id
	claims["exp"] = authClaims.Expiration
}

func (authClaims *AuthClaims) Parse(token *jwt.Token) error {
	claims := token.Claims.(jwt.MapClaims)
	authClaims.Id = claims["id"].(string)
	authClaims.Expiration = claims["exp"].(time.Time)

	if authClaims.Id == "" {
		return jwt.ValidationError{}
	}

	return nil
}
