package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"playgrounds.com/user"
)

func main() {

	os.Setenv("MongoUrl", "mongodb://myUser:myPassword@localhost:27017/")
	os.Setenv("MongoDBName", "PhotoManager")

	r := gin.Default()
	api := r.Group("/api")
	userApi := api.Group("/user")
	user.Setup(os.Getenv("MongoUrl"), os.Getenv("MongoDBName"))
	user.HandleRoutes(userApi)
	r.Run(":8000")
}
