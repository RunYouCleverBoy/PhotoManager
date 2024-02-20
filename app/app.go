package main

import (
	"github.com/gin-gonic/gin"
	"playgrounds.com/auth"
	"playgrounds.com/database"
	"playgrounds.com/environment"
	"playgrounds.com/photos"
	"playgrounds.com/user"
	"playgrounds.com/utils"
)

func main() {

	var fakeEnv environment.Environment = environment.Environment{
		DatabaseURL:  "mongodb://localhost:27017/",
		DatabaseName: "PhotoManager",
		JWTSecret:    []byte("Sittin' in the stand of the sports arena, waiting for the show to begin Red lights, green lights, strawberry wine, a good friend of mine, follows the stars Venus and Mars are alright tonight"),
	}

	environment.ApplyEnvironment(&fakeEnv)

	env := environment.NewFromEnv()
	db, err := database.NewDb(env.DatabaseURL, env.DatabaseName)
	authMiddleware := utils.AuthMiddleware(&env.JWTSecret)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	api := r.Group("/api")
	userApi := api.Group("/user")
	user.Setup(db.UserCollection())
	user.HandleRoutes(userApi, authMiddleware)

	photosApi := api.Group("/photos")
	photos.Setup(db.PhotosCollection())
	photos.HandleRoutes(photosApi)

	authApi := api.Group("/auth")
	auth.Setup(env)
	auth.HandleRoutes(authApi, authMiddleware)
	r.Run(":8000")
}
