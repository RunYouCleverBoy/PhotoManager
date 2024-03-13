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
	log.Printf("\n" + env.String())

	db, err := database.NewDb(env.DatabaseURL, env.DatabaseName)
	if err != nil {
		panic(err)
	}

	authMiddleware := auth.AuthMiddleware(&env.JWTSecret)

	r := gin.Default()
	api := r.Group("/api")
	userApi := api.Group("/user")
	user.Setup(db.UserCollection())
	user.HandleRoutes(userApi, authMiddleware)
	log.Print("Server: User routes initialized")

	photosApi := api.Group("/photos")
	photos.Setup(db.PhotosCollection())
	photos.HandleRoutes(photosApi, authMiddleware)
	log.Print("Server: Photos routes initialized")

	albumsApi := api.Group("/albums")
	photoalbums.Setup(db.AlbumsCollection(), db.PhotosCollection())
	photoalbums.HandleRoutes(albumsApi, authMiddleware)
	log.Print("Server: Albums routes initialized")

	authApi := api.Group("/auth")
	auth.Setup(env)
	auth.HandleRoutes(authApi, authMiddleware)
	log.Print("Server: Auth routes initialized")

	r.Run(fmt.Sprintf(":%s", env.Port))
}
