package photos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"playgrounds.com/models"
)

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
