package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthClaims struct {
	Id         *primitive.ObjectID
	Expiration *time.Time
}

func (authClaims *AuthClaims) ApplyClaims(token *jwt.Token) {
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = authClaims.Id.Hex()
	claims["exp"] = authClaims.Expiration
}

func (authClaims *AuthClaims) Parse(token *jwt.Token) error {
	claims := token.Claims.(jwt.MapClaims)
	idStr := claims["id"].(string)
	if idStr == "" {
		return jwt.ValidationError{}
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return jwt.ValidationError{}
	}
	authClaims.Id = &id
	expiration := claims["exp"].(time.Time)
	authClaims.Expiration = &expiration

	if authClaims.Id.IsZero() {
		return jwt.ValidationError{}
	}

	return nil
}
