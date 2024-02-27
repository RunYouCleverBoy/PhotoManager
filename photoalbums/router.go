package photoalbums

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.Use(authMiddleware)
	router.GET("/myalbums", GetMyAlbums)
	router.GET("/albums/:id", RequireAlbumVisibility("id"), GetAlbum)
	router.POST("/albums/create", CreateAlbum)

	router.Use(RequireAlbumOwner("id"))
	router.POST("/albums/:id/addVisibility", AddOrRemoveAlbumVisibility)

	router.DELETE("/albums/:id", DeleteAlbum)

}
