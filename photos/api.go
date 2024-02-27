package photos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/models"
)

type CreatePhotoApi struct {
	Id       primitive.ObjectID    `json:"id"`
	IsPublic bool                  `json:"isPublic"`
	Metadata *models.PhotoMetadata `json:"metadata,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Comments []models.Comments     `json:"comments,omitempty"`
}

type EditPhotoApi struct {
	IsPublic *bool                 `json:"isPublic,omitempty"`
	Metadata *models.PhotoMetadata `json:"metadata,omitempty"`
	WorkFlow *models.WorkFlow      `json:"workflow,omitempty"`
}

type AlbumAndPhotoRequestBody struct {
	AlbumId string   `json:"albumId"`
	PhotoId []string `json:"photoId"`
}
