package photoalbums

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.Use(authMiddleware)
	router.GET("/mine", GetMyAlbums)
	router.GET("/search", SearchAlbum)
	router.GET("/:id", RequireAlbumVisibility("id"), GetAlbum)
	router.POST("/create", CreateAlbum)

	router.Use(RequireAlbumOwner("id"))
	router.POST("/:id/addvisibility", AddOrRemoveAlbumVisibility)
	router.POST("/:id/addRemovephotos", AddAndRemovePhotosToAlbum)

	router.DELETE("/:id", DeleteAlbum)
}
