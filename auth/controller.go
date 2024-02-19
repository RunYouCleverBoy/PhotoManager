package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"playgrounds.com/database"
	"playgrounds.com/environment"
	"playgrounds.com/user"
	"playgrounds.com/utils"
)

var signFunc func(token *jwt.Token) (string, error) = nil

func Setup(env *environment.Environment) {
	signFunc = func(token *jwt.Token) (string, error) {
		return token.SignedString([]byte(env.JWTSecret))
	}
}

func login(ctx *gin.Context) {
	requestBody := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	ctx.Bind(&requestBody)
	userObj, err := user.GetUserByEmail(requestBody.Username)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(requestBody.Password)) != nil {
		ctx.JSON(401, gin.H{"message": "invalid credentials"})
		return
	}

	token, authClaims, err := createToken(userObj)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	expiration := authClaims.Expiration.Unix()
	userObj, err = user.UpdateCredentials(&userObj.ID, token, expiration)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"token": userObj})
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
	token, authClaims, err := createToken(&userObj)
	if err != nil {
		respondToError(ctx, err)
		return
	}

	expiration := authClaims.Expiration.Unix()
	userObj.TokenExpiry = &expiration
	userObj.Token = token

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
	objIdValue, exists := ctx.Get(utils.CallingUserIdContextKey)
	if !exists {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": "Invalid JWT"})
		return
	}

	id := objIdValue.(primitive.ObjectID)
	user.UpdateCredentials(&id, nil, 0)
	ctx.JSON(200, gin.H{"message": "logout"})
}

func respondToError(ctx *gin.Context, err error) {
	if err == database.ErrInvalidId {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
	} else {
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
	}
}

func createToken(user *user.User) (*string, *utils.AuthClaims, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	authClaims := utils.AuthClaims{}
	authClaims.Id = &user.ID
	expiration := time.Now().Add(time.Hour * 24 * 90)
	authClaims.Expiration = &expiration
	authClaims.ApplyClaims(token)
	tokenString, err := signFunc(token)
	if err != nil {
		return nil, nil, err
	}
	return &tokenString, &authClaims, nil
}
