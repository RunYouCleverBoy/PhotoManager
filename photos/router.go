package photos

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup) {
	router.GET("/", GetAllMyPhotos)
	router.GET("/:id", GetPhoto)
	router.POST("/", CreatePhoto)
	router.PUT("/:id", UpdatePhoto)
	router.DELETE("/:id", DeletePhoto)
}
