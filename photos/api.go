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

func (req *CreatePhotoApi) toPhotoModel(ownerId primitive.ObjectID) models.PhotoModel {
	objectId := primitive.NewObjectID()
	return models.PhotoModel{
		ID:        objectId,
		IsPublic:  false,
		Owner:     ownerId,
		VisibleTo: []primitive.ObjectID{},
		Metadata:  *req.Metadata,
		WorkFlow: models.WorkFlow{
			UpvoteGrade:   0,
			WorkflowStage: models.WorkflowStageAlbum,
		},
		Ancestor: objectId, // A footage is its own ancestor
		Tags:     req.Tags,
		Comments: req.Comments,
	}
}

type AlbumAndPhotoRequestBody struct {
	AlbumId string   `json:"albumId"`
	PhotoId []string `json:"photoId"`
}

type AddOrRemovePhotosRequestBody struct {
	AddPhotos    []primitive.ObjectID `json:"addPhotos"`
	RemovePhotos []primitive.ObjectID `json:"removePhotos"`
}

type AddOrRemoveVisibilityRequestBody struct {
	AddVisibleTo    []primitive.ObjectID `json:"addVisibleTo"`
	RemoveVisibleTo []primitive.ObjectID `json:"removeVisibleTo"`
}
