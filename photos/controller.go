package photos

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/database"
	"playgrounds.com/models"
	"playgrounds.com/utils"
)

var db *database.PhotosCollection

var (
	ErrorNotLoggedIn     = errors.New("not logged in")
	ErrorInvalidObjectId = errors.New("invalid object id")
)

func Setup(collection *database.PhotosCollection, albumsCollection *database.AlbumCollection) {
	db = collection
	albums = albumsCollection
}

func GetAllMyPhotos(c *gin.Context) {
	userId := utils.CollectIdFromAuthentication(c)

	indexRange := getPagingQueryArgs(c)

	if photos, error := db.GetPhotosByUserId(userId, indexRange); error != nil {
		c.JSON(500, gin.H{"message": "error", "error": error.Error()})
	} else {
		c.JSON(200, photos)
	}
}

func SearchPhotos(c *gin.Context) {

	searchOptions := models.PhotoSearchOptions{}
	if err := c.BindJSON(&searchOptions); err != nil {
		c.JSON(400, gin.H{"message": "invalid search options", "error": err.Error()})
		return
	}

	callerId := utils.CollectIdFromAuthentication(c)
	if callerId == nil {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}

	indexRange := getPagingQueryArgs(c)

	if photos, error := db.GetPhotos(callerId, &searchOptions, indexRange); error != nil {
		c.JSON(500, gin.H{"message": "error", "error": error.Error()})
	} else {
		c.JSON(200, photos)
	}
}

func GetPhoto(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id " + idStr, "error": err.Error()})
		return
	}

	if photo, e := db.GetPhotoById(&id); e != nil {
		c.JSON(500, gin.H{"message": "error", "error": e.Error()})
	} else if photo == nil {
		c.JSON(404, gin.H{"message": "not found", "id": idStr})
	} else {
		c.JSON(200, photo)
	}
}

func CreatePhoto(c *gin.Context) {
	createJson := CreatePhotoApi{}
	if error := c.BindJSON(&createJson); error != nil {
		c.JSON(400, gin.H{"message": "invalid photo", "error": error.Error()})
		return
	}

	photo := createJson.toPhotoModel(*utils.CollectIdFromAuthentication(c))
	currentUserId := utils.CollectIdFromAuthentication(c)
	if photo, error := db.CreatePhoto(currentUserId, photo); error != nil {
		c.JSON(500, gin.H{"message": "error", "error": error.Error()})
	} else {
		c.JSON(201, photo)
	}
}

func AddComment(c *gin.Context) {
	comment := models.Comments{}
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"message": "invalid photo", "error": err.Error()})
		return
	}

	photoIdStr := c.Param("id")
	photoId, err := primitive.ObjectIDFromHex(photoIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid photo ID", "error": err.Error()})
		return
	}

	if comment, e := db.AddCommentToPhoto(&photoId, &comment); e != nil {
		c.JSON(500, gin.H{"message": "error", "error": e.Error()})
	} else {
		c.JSON(200, comment)
	}
}

func AddTag(c *gin.Context) {
	tag := struct {
		Tag string `json:"tag"`
	}{}
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(400, gin.H{"message": "invalid photo", "error": err.Error()})
		return
	}

	photoIdStr := c.Param("id")
	photoId, err := primitive.ObjectIDFromHex(photoIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid photo ID", "error": err.Error()})
		return
	}

	if tag, e := db.AddTagToPhoto(&photoId, tag.Tag); e != nil {
		c.JSON(500, gin.H{"message": "error", "error": e.Error()})
	} else {
		c.JSON(200, tag)
	}
}

func UpdatePhoto(c *gin.Context) {
	// TODO
}

func DeletePhoto(c *gin.Context) {
	// TODO
}

func GetPublicPhoto(c *gin.Context) {
	id := c.Param("id")
	photo, err := getPhotoById(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid photo ID", "error": err.Error()})
		return
	}

	if !photo.IsPublic {
		c.JSON(403, gin.H{"message": "photo not public"})
		return
	}

	c.JSON(200, photo)
}

func RequireVisibility(photoIdParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoId := c.Param(photoIdParam)
		if photoId == "" {
			c.AbortWithError(400, errors.New("missing photo id"))
			return
		}

		photo, err := getPhotoById(photoId)
		if err != nil {
			abortWithError(c, err)
			return
		}

		if !photo.IsPublic {
			c.AbortWithError(403, errors.New("photo not public"))
			return
		}

		c.Next()
	}
}

func RequireOwner(photoIdParam string) gin.HandlerFunc {
	return func(c *gin.Context) {
		photoId := c.Param(photoIdParam)
		if photoId == "" {
			c.AbortWithError(400, errors.New("missing photo id"))
			return
		}

		userId := utils.CollectIdFromAuthentication(c)
		if userId == nil {
			c.AbortWithError(401, ErrorNotLoggedIn)
			return
		}

		photo, err := getPhotoById(photoId)
		if err != nil {
			abortWithError(c, err)
			return
		}

		if photo.Owner.Hex() != userId.Hex() {
			c.AbortWithError(403, errors.New("not owner of photo"))
			return
		}

		c.Next()
	}
}

// router.DELETE("/albums/:id", RequireAlbumOwner("id"), DeleteAlbum)
func AddAndRemovePhotosToAlbum(c *gin.Context) {
	albumIdStr := c.Param("id")
	albumId, err := primitive.ObjectIDFromHex(albumIdStr)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid album ID", "error": err.Error()})
		return
	}

	requestBody := AddOrRemovePhotosRequestBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"message": "invalid request body", "error": err.Error()})
		return
	}

	err = db.AddOrRemoveAlbumFromManyPhotos(&albumId, &requestBody.AddPhotos, &requestBody.RemovePhotos)
	if err != nil {
		c.JSON(500, gin.H{"message": "error", "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

func getPagingQueryArgs(c *gin.Context) *utils.IntRange[int] {
	query := c.Request.URL.Query()
	var err error
	var startIndex, pageSize int
	if startIndex, err = strconv.Atoi(query.Get("startindex")); err != nil {
		startIndex = 0
	}
	if pageSize, err = strconv.Atoi(query.Get("limitindex")); err != nil {
		pageSize = 100
	}

	return utils.NewIntRange(startIndex, pageSize)
}

func getPhotoById(id string) (*models.PhotoModel, error) {
	photoId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrorInvalidObjectId
	}

	photo, err := db.GetPhotoById(&photoId)
	if err != nil {
		return nil, errors.New("error getting photo: " + err.Error())
	}

	return photo, nil
}

func abortWithError(c *gin.Context, err error) {
	status, errMsg := resolveError(c, err)
	c.AbortWithError(status, errors.New(errMsg+" "+err.Error()))
}

func resolveError(c *gin.Context, err error) (int, string) {
	errCode, errMsg := 500, "error"
	switch err {
	case ErrorNotLoggedIn:
		errCode, errMsg = 401, "not logged in"
	case ErrorInvalidObjectId:
		errCode, errMsg = 400, "invalid object id"
	default:
		errCode, errMsg = 500, err.Error()
	}
	return errCode, errMsg
}
