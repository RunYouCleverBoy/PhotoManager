package photos

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.GET("/my", authMiddleware, GetAllMyPhotos)
	router.POST("/search", authMiddleware, SearchPhotos)
	router.GET("/:id", authMiddleware, RequireVisibility("id"), GetPhoto)
	router.GET("/public/:id", GetPublicPhoto)
	router.POST("/create", authMiddleware, CreatePhoto)
	router.POST("/addcomment/:id", authMiddleware, RequireVisibility("id"), AddComment)
	router.POST("/addtag/:id", authMiddleware, RequireOwner("id"), AddTag)
	router.DELETE("/:id", DeletePhoto)

	router.Use(authMiddleware)
	router.GET("/myalbums", GetMyAlbums)
	router.GET("/albums/:id", RequireAlbumVisibility("id"), GetAlbum)
	router.POST("/albums/create", CreateAlbum)

	router.Use(RequireVisibility("id"))
	router.POST("/albums/:id/addRemovephotos", RequireAlbumOwner("id"), AddAndRemovePhotosToAlbum)
	router.POST("/albums/:id/addVisibility", RequireAlbumOwner("id"), AddOrRemoveAlbumVisibility)

	router.DELETE("/albums/:id", RequireAlbumOwner("id"), DeleteAlbum)
}
