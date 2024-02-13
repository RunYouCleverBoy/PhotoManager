package main

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/user"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")
	userApi := api.Group("/user")
	user.HandleRoutes(userApi)
	r.Run(":8000")
}
