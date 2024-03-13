package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"playgrounds.com/auth"
	"playgrounds.com/database"
	"playgrounds.com/environment"
	"playgrounds.com/photoalbums"
	"playgrounds.com/photos"
	"playgrounds.com/user"
)

func main() {

	env := environment.NewFromEnv(&environment.DefaultEnvironment)
	log.Print("\n" + env.String())

	db, err := database.NewDb(env.DatabaseURL, env.DatabaseName)
	authMiddleware := auth.AuthMiddleware(&env.JWTSecret)

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
	photos.HandleRoutes(photosApi, authMiddleware)

	albumsApi := api.Group("/albums")
	photoalbums.Setup(db.AlbumsCollection(), db.PhotosCollection())
	photoalbums.HandleRoutes(albumsApi, authMiddleware)

	authApi := api.Group("/auth")
	auth.Setup(env)
	auth.HandleRoutes(authApi, authMiddleware)

	r.Run(fmt.Sprintf(":%s", env.Port))
}
