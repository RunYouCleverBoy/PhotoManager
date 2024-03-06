package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"playgrounds.com/database"
	"playgrounds.com/environment"
	"playgrounds.com/user"
	"playgrounds.com/utils"
)

var signFunc func(token *jwt.Token) (string, error) = nil
var verifyFunc func(tokenStr *string) (*jwt.Token, error) = nil
var refreshTokenExpiration, tokenExpiration *time.Duration

func Setup(env *environment.Environment) {
	signFunc = func(token *jwt.Token) (string, error) {
		return token.SignedString(env.JWTSecret)
	}

	verifyFunc = func(tokenStr *string) (*jwt.Token, error) {
		token, err := jwt.Parse(*tokenStr, func(token *jwt.Token) (interface{}, error) {
			return env.JWTSecret, nil
		})

		return token, err
	}

	tokenExpiration = &env.TokenExpiration
	refreshTokenExpiration = &env.RefreshTokenExpiration
}

func login(ctx *gin.Context) {
	requestBody := LoginRequest{}
	ctx.Bind(&requestBody)
	if requestBody.Email == "" || requestBody.Password == "" {
		ctx.JSON(400, gin.H{"message": "invalid request", "error": "missing email or password"})
		return
	}

	userObj, err := user.GetUserByEmail(requestBody.Email)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(requestBody.Password)) != nil {
		ctx.JSON(401, gin.H{"message": "invalid credentials"})
		return
	}

	token, _, err := createToken(userObj, tokenExpiration)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	refreshToken, _, err := createToken(userObj, refreshTokenExpiration)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	userObj, err = user.UpdateCredentials(&userObj.ID, &userObj.Password, token, refreshToken)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"user": userObj})
}

func loginWithToken(ctx *gin.Context) {
	requestBody := LoginWithTokenRequest{}
	if err := ctx.Bind(&requestBody); err != nil || requestBody.RefreshToken == "" || requestBody.OldToken == "" {
		ctx.JSON(400, gin.H{"message": "invalid request", "error": "missing refresh token"})
		return
	}

	oldToken, err := verifyFunc(&requestBody.OldToken)
	oldTokenValid := false
	switch {
	case oldToken.Valid || err != nil:
		oldTokenValid = true
	case errors.Is(err, jwt.ErrTokenExpired):
		oldTokenValid = true
	default:
		oldTokenValid = false
		return
	}

	if !oldTokenValid {
		ctx.JSON(401, gin.H{"message": "invalid old token", "error": err.Error()})
		return
	}

	refreshToken, err := verifyFunc(&requestBody.RefreshToken)
	if err != nil {
		ctx.JSON(401, gin.H{"message": "invalid refresh token", "error": err.Error()})
		return
	}

	claims := AuthClaims{}
	if err := claims.Parse(refreshToken); err != nil {
		ctx.JSON(401, gin.H{"message": "invalid refresh token", "error": err.Error()})
		return
	}

	userId := claims.Id
	userObj, err := user.GetUserById(userId)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	tokenStr, refreshTokenStr, err := createTokens(userObj)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	userObj, err = user.UpdateCredentials(userId, nil, tokenStr, refreshTokenStr)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"user": userObj})
}

func register(ctx *gin.Context) {
	userObj := user.User{}
	ctx.Bind(&userObj)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userObj.Password), bcrypt.DefaultCost)
	if err != nil {
		respondToError(ctx, err)
		return
	}
	userObj.Password = string(hashedPassword)
	userObj.ID = primitive.NewObjectID()
	token, refreshToken, err := createTokens(&userObj)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	userObj.Token = token
	userObj.RefreshToken = refreshToken

	// Create the user
	responseUser, createErr := user.CreateUserByUserObject(&userObj)
	if createErr != nil {
		respondToError(ctx, createErr)
		return
	}

	responseUser.Password = string(hashedPassword)
	ctx.JSON(201, gin.H{"user": responseUser})
}

func logout(ctx *gin.Context) {
	userId := utils.CollectIdFromAuthentication(ctx)
	if userId == nil {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": "Invalid JWT"})
		return
	}

	emptyString := ""
	user.UpdateCredentials(userId, nil, &emptyString, &emptyString)
	ctx.JSON(200, gin.H{"message": "logout"})
}

func respondToError(ctx *gin.Context, err error) {
	if err == database.ErrInvalidId {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
	} else {
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
	}
}

func createTokens(userObj *user.User) (*string, *string, error) {
	token, _, err := createToken(userObj, tokenExpiration)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, _, err := createToken(userObj, refreshTokenExpiration)
	if err != nil {
		return nil, nil, err
	}

	return token, refreshToken, nil
}

func createToken(user *user.User, expirationDuration *time.Duration) (*string, *AuthClaims, error) {
	authClaims := AuthClaims{}
	authClaims.Id = &user.ID
	expiration := time.Now().Add(*expirationDuration)
	authClaims.Expiration = &expiration
	token := authClaims.NewUnsignedToken(time.Second * 30)
	tokenString, err := signFunc(token)
	if err != nil {
		return nil, nil, err
	}

	return &tokenString, &authClaims, nil
}
