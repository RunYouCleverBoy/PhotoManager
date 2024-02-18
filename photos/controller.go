package photos

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/database"
	"playgrounds.com/environment"
)

func GetAllPhotos(c *gin.Context) {
	// TODO
}

func GetPhoto(c *gin.Context) {
	// TODO
}

func CreatePhoto(c *gin.Context) {
	// TODO
}

func UpdatePhoto(c *gin.Context) {
	// TODO
}

func DeletePhoto(c *gin.Context) {
	// TODO
}

func Setup(env environment.Environment, collection *database.PhotosCollection) {
	// TODO
}

func restrictTo(role string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
