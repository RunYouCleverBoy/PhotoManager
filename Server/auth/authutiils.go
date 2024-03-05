package auth

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"playgrounds.com/utils"
)

var (
	ErrorMissingToken = errors.New("missing token")
	ErrorBadHeader    = errors.New("bad header value given")
	ErrorBadToken     = errors.New("bad token")
)

const (
	CallingUserIdContextKey = utils.CallingUserIdContextKey
	authClaimsContextKey    = "authClaims"
	jwtTokenContextKey      = "jwtToken"
)

func AuthMiddleware(jwtSecret *[]byte) gin.HandlerFunc {
	validateJWT := func(token string, secret []byte) (*jwt.Token, error) {
		return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
	}
	return func(ctx *gin.Context) {
		token, err := extractBearerToken(ctx)
		if err != nil {
			ctx.AbortWithError(401, err)
			return
		}

		jwtToken, err := validateJWT(token, *jwtSecret)
		if err != nil {
			ctx.AbortWithError(401, err)
			return
		}

		authClaims := AuthClaims{}
		fieldErr := authClaims.Parse(jwtToken)
		if fieldErr != nil {
			ctx.AbortWithError(401, fieldErr)
			return
		}

		ctx.Set(CallingUserIdContextKey, *authClaims.Id)
		ctx.Set(authClaimsContextKey, authClaims)
		ctx.Set(jwtTokenContextKey, jwtToken)
		ctx.Next()
	}
}

func extractBearerToken(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return "", ErrorMissingToken
	}

	regexSlice := regexp.MustCompile("Bearer (.+)").FindStringSubmatch(header)
	if regexSlice == nil || len(regexSlice) != 2 {
		return "", ErrorBadHeader
	}

	return regexSlice[1], nil
}
