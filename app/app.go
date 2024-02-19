package main

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/auth"
	"playgrounds.com/database"
	"playgrounds.com/environment"
	"playgrounds.com/user"
	"playgrounds.com/utils"
)

func main() {

	var fakeEnv environment.Environment = environment.Environment{
		DatabaseURL:  "mongodb://localhost:27017/",
		DatabaseName: "PhotoManager",
		JWTSecret:    "Bolshoi bolshoi secret da",
	}

	environment.ApplyEnvironment(&fakeEnv)

	env := environment.NewFromEnv()
	db, err := database.NewDb(env.DatabaseURL, env.DatabaseName)
	authMiddleware := utils.AuthMiddleware(env.JWTSecret)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	api := r.Group("/api")
	userApi := api.Group("/user")
	user.Setup(db.UserCollection())
	user.HandleRoutes(userApi, authMiddleware)

	// photosApi := api.Group("/photos")
	// photosApi.Setup(env, db.PhotoCollection())
	// photosApi.HandleRoutes(photosApi)

	authApi := api.Group("/auth")
	auth.Setup(env)
	auth.HandleRoutes(authApi, authMiddleware)
	r.Run(":8000")
}
