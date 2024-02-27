package photos

import (
	"playgrounds.com/models"
)

type CreatePhotoRequest struct {
	IsPublic bool                  `json:"isPublic"`
	Metadata *models.PhotoMetadata `json:"metadata,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Comments []models.Comments     `json:"comments,omitempty"`
}

type CreatePhotoResponse struct {
	UploadUrl   string             `json:"uploadUrl"`
	PhotoRecord *models.PhotoModel `json:"photoRecord"`
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
