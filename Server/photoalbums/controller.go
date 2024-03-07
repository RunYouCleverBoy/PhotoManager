package photoalbums

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/database"
	"playgrounds.com/models"
	"playgrounds.com/user"
	"playgrounds.com/utils"
)

const (
	AlbumContextKey = "album"
)

var albums *database.AlbumCollection
var photos *database.PhotosCollection

func Setup(albumCollection *database.AlbumCollection, photoCollection *database.PhotosCollection) {
	albums = albumCollection
	photos = photoCollection
}

func GetMyAlbums(c *gin.Context) {
	userId := utils.CollectIdFromAuthentication(c)
	albums, err := albums.GetAlbumsBy(&AlbumSearchCriteria{OwnerID: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func GetAlbum(c *gin.Context) {
	if album, ok := c.Get(AlbumContextKey); ok {
		c.JSON(http.StatusOK, album)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
	}
}

func CreateAlbum(c *gin.Context) {
	userId := utils.CollectIdFromAuthentication(c)
	var album models.PhotoAlbum
	if err := c.BindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album.ID = primitive.NewObjectID()
	album.Owner = *userId
	if err := albums.CreateAlbum(&album); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, album)
}

func AddOrRemoveAlbumVisibility(c *gin.Context) {
	albumIdStr := c.Param("albumId")
	albumId, err := primitive.ObjectIDFromHex(albumIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid album id"})
		return
	}

	reqBody := AddOrRemoveVisibilityRequestBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allUsers := append(reqBody.AddVisibleTo, reqBody.RemoveVisibleTo...)
	if allExist, err := user.VerifyUsersExist(&allUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if !allExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "one or more users do not exist"})
		return
	}

	if err := albums.AddOrRemoveAlbumVisibility(&albumId, &reqBody.AddVisibleTo, &reqBody.RemoveVisibleTo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func AddAndRemovePhotosToAlbum(c *gin.Context) {
	albumIdStr := c.Param("id")
	albumId, err := primitive.ObjectIDFromHex(albumIdStr)
	userId := utils.CollectIdFromAuthentication(c)

	if err != nil {
		c.JSON(400, gin.H{"message": "invalid album ID", "error": err.Error()})
		return
	}

	requestBody := AddOrRemovePhotosRequestBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	allPhotos := append(requestBody.AddPhotos, requestBody.RemovePhotos...)
	err = photos.VerifyVisibilityForAllPhotos(&allPhotos, userId)
	if err != nil {
		c.JSON(403, gin.H{"message": "not all photos are accessible", "error": err.Error()})
		return
	}

	err = photos.AddOrRemoveAlbumFromManyPhotos(&albumId, &requestBody.AddPhotos, &requestBody.RemovePhotos)
	if err != nil {
		c.JSON(500, gin.H{"message": "error", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func DeleteAlbum(c *gin.Context) {
	albumIdStr := c.Param("id")
	albumId, err := primitive.ObjectIDFromHex(albumIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid album id"})
		return
	}

	err = albums.DeleteAlbum(&albumId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func RequireAlbumOwner(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		albumId := c.Param(paramName)
		album, err := albums.GetAlbumById(&albumId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			c.Abort()
			return
		}

		userId := utils.CollectIdFromAuthentication(c)
		if album.Owner != *userId {
			c.JSON(http.StatusForbidden, gin.H{"error": "user not album owner"})
			c.Abort()
			return
		}

		c.Set(AlbumContextKey, album)
		c.Next()
	}
}

func RequireAlbumVisibility(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		albumId := c.Param(paramName)
		album, err := albums.GetAlbumById(&albumId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			c.Abort()
			return
		}

		userId := utils.CollectIdFromAuthentication(c)
		if !slices.Contains(album.VisibleTo, *userId) && !(album.Owner == *userId) {
			c.JSON(http.StatusForbidden, gin.H{"error": "album not visible to user"})
			c.Abort()
			return
		}

		c.Set(AlbumContextKey, album)
		c.Next()
	}
}
