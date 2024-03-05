package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthClaims struct {
	Id         *primitive.ObjectID
	Expiration *time.Time
}

func (authClaims *AuthClaims) NewUnsignedToken(issueLeeway time.Duration) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  authClaims.Id.Hex(),
		"exp": authClaims.Expiration.Unix(),
		"iat": time.Now().Add(-issueLeeway).Unix(),
	})
	return token
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
	expirationUnix := int64(claims["exp"].(float64))
	expiration := time.Unix(expirationUnix, 0)
	authClaims.Expiration = &expiration

	if authClaims.Id.IsZero() {
		return jwt.ValidationError{}
	}

	return nil
}
