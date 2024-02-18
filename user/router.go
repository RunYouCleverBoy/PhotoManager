package user

import (
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	router.Use(authMiddleware, getCurrentUser)

	router.GET("/", restrictTo(RoleAdmin), GetAllUsers)
	router.GET("/:id", restrictTo(RoleAdmin), GetUser)
	router.POST("/", restrictTo(RoleAdmin), CreateUser)
	router.PUT("/:id", restrictTo(RoleAdmin), omitFields, UpdateUser)
	router.DELETE("/:id", restrictTo(RoleAdmin), DeleteUser)

	router.GET("/me", restrictTo(RoleUser), selfService)
	router.PUT("/me", restrictTo(RoleUser), selfService)
	router.DELETE("/me", restrictTo(RoleUser), selfService)
}
