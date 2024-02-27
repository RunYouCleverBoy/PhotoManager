package photos

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/photoalbums"
)

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.GET("/my", authMiddleware, GetAllMyPhotos)
	router.POST("/search", authMiddleware, SearchPhotos)
	router.GET("/:id", authMiddleware, RequireVisibility("id"), GetPhoto)
	router.GET("/public/:id", GetPublicPhoto)
	router.POST("/create", authMiddleware, CreatePhoto)
	router.POST("/addcomment/:id", authMiddleware, RequireVisibility("id"), AddComment)
	router.POST("/addtag/:id", authMiddleware, RequireOwner("id"), AddTag)
	router.DELETE("/:id", DeletePhoto)

	router.POST("/albums/:id/addRemovephotos", photoalbums.RequireAlbumOwner("id"), AddAndRemovePhotosToAlbum)
}
