package photos

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/database"
	"playgrounds.com/utils"
)

var db *database.PhotosCollection

func Setup(collection *database.PhotosCollection) {
	db = collection
}

func GetAllMyPhotos(c *gin.Context) {
	userId := utils.CollectDataFromAuthentication(c)
	photos, err := db.GetPhotoById()
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

func restrictTo(role string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
