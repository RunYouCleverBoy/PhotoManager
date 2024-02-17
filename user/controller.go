package user

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/database"
	"playgrounds.com/models"
)

type User = models.User

var db *database.Database

func Setup(mongoUrl string, mongoDBName string) {
	err := error(nil)
	db, err = database.NewDb(mongoUrl, mongoDBName)
	if err != nil {
		panic(err)
	}
}

func GetAllUsers(ctx *gin.Context) {
	data, err := db.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
		return
	}
	ctx.JSON(200, data)
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := db.Get(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
	} else if user == nil {
		ctx.JSON(404, gin.H{"message": "not found", "id": id})
	} else {
		ctx.JSON(200, user)
	}
}

func CreateUser(ctx *gin.Context) {
	user := User{}
	ctx.Bind(&user)
	result, err := db.Create(user)
	if err != nil {
		respondToError(ctx, err)
		return
	}
	ctx.JSON(201, result)
}

func UpdateUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(User)
	id := ctx.Param("id")
	result, err := db.Update(id, user)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "error", "error": "error updating user"})
		return
	}
	ctx.JSON(200, &result)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := db.Delete(id); err != nil {
		ctx.JSON(500, gin.H{"message": "error", "error": "error deleting user"})
		return
	}
	ctx.JSON(200, gin.H{"message": "delete", "id": id})
}

func OmitFields(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": "id is required"})
		return
	}

	user := User{}
	ctx.Bind(&user)

	user.ID = ""
	user.Token = nil
	user.TokenExpiry = nil

	ctx.Set("user", user)
	ctx.Set("id", id)

	ctx.Next()
}

func respondToError(ctx *gin.Context, err error) {
	switch err {
	case nil:
		return
	case database.ErrInvalidId:
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
	case database.ErrInvalidUser:
		ctx.JSON(400, gin.H{"message": "invalid user", "error": err.Error()})
	default:
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
	}
}
