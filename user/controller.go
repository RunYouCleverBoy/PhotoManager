package user

import (
	"slices"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/database"
	"playgrounds.com/models"
	"playgrounds.com/utils"
)

type User = models.User

const (
	RoleAdmin Role = models.RoleAdmin
	RoleUser  Role = models.RoleUser
)

const CallingUserContextKey = "CallingUser"
const SubjectUserContextKey = "SubjectUser"

type Role = models.Role

var db *database.UserCollection

func Setup(collection *database.UserCollection) {
	db = collection
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
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
		return
	}

	user, err := db.Get(&id)
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
	result, err := CreateUserByUserObject(&user)
	if err != nil {
		respondToError(ctx, err)
		return
	}
	ctx.JSON(201, result)
}

func UpdateUser(ctx *gin.Context) {
	var user User
	if data, exists := ctx.Get(SubjectUserContextKey); exists {
		user = data.(User)
	} else {
		user = User{}
		ctx.Bind(&user)
	}

	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
		return
	}

	result, err := db.Update(&id, &user)
	if err != nil {
		respondToError(ctx, err)
		return
	}
	ctx.JSON(200, &result)
}

func DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": err.Error()})
		return
	}

	if err := db.Delete(&id); err != nil {
		respondToError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "delete", "id": id})
}

func restrictTo(roles ...Role) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		requestingUser := ctx.MustGet(CallingUserContextKey).(User)
		if *requestingUser.Role == RoleAdmin {
			ctx.Next()
		}

		isInRole := slices.Contains(roles, *requestingUser.Role)
		if !isInRole {
			ctx.JSON(403, gin.H{"message": "forbidden", "error": "insufficient permissions"})
			ctx.Abort()
		}

		ctx.Next()
	}
}

func omitFields(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"message": "invalid id", "error": "id is required"})
		return
	}

	user := User{}
	ctx.Bind(&user)

	user.ID = primitive.ObjectID{}
	user.Token = nil
	user.TokenExpiry = nil
	user.Role = nil

	ctx.Set(SubjectUserContextKey, user)
	ctx.Set("id", id)

	ctx.Next()
}

func selfService(ctx *gin.Context) {
	user := ctx.MustGet(CallingUserContextKey).(User)
	id := &user.ID
	switch ctx.Request.Method {
	case "GET":
		ctx.JSON(200, user)
	case "PUT":
		user := User{}
		ctx.Bind(&user)
		result, err := db.Update(id, &user)
		if err != nil {
			respondToError(ctx, err)
			return
		}
		ctx.JSON(200, result)
	case "DELETE":
		if err := db.Delete(id); err != nil {
			respondToError(ctx, err)
			return
		}
		ctx.JSON(200, gin.H{"message": "delete", "id": id})
	}
}

func getCurrentUser(ctx *gin.Context) {
	id := ctx.MustGet(CallingUserContextKey).(primitive.ObjectID)
	user, err := db.Get(&id)
	if err != nil {
		respondToError(ctx, err)
		ctx.Abort()
		return
	}
	ctx.Set(CallingUserContextKey, user)
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
	case database.ErrNoDocuments:
		ctx.JSON(404, gin.H{"message": "not found", "error": err.Error()})
	case utils.ErrorBadHeader, utils.ErrorMissingToken, utils.ErrorBadToken:
		ctx.JSON(401, gin.H{"message": "unauthorized", "error": err.Error()})
	default:
		ctx.JSON(500, gin.H{"message": "error", "error": err.Error()})
	}
}
