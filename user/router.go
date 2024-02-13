package user

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup) {
	router.GET("/", GetAllUsers)
	router.GET("/:id", GetUser)
	router.POST("/", CreateUser)
	router.PUT("/:id", UpdateUser)
	router.DELETE("/:id", DeleteUser)
}
