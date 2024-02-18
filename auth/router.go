package auth

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/register", register)
	router.POST("/logout", logout)
}
