package utils

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CallingUserIdContextKey = "callingUserId"

func CollectIdFromAuthentication(ctx *gin.Context) (id *primitive.ObjectID) {
	if id, exists := ctx.Get(CallingUserIdContextKey); exists {
		userID := id.(primitive.ObjectID)
		return &userID
	} else {
		return nil
	}
}
