package auth

import "github.com/gin-gonic/gin"

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.POST("/login", login)
	router.POST("/register", register)
	router.POST("/logout", authMiddleware, logout)
}
