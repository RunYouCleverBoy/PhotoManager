package photoalbums

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddOrRemoveVisibilityRequestBody struct {
	AddVisibleTo    []primitive.ObjectID `json:"addVisibleTo"`
	RemoveVisibleTo []primitive.ObjectID `json:"removeVisibleTo"`
}

type AddOrRemovePhotosRequestBody struct {
	AddPhotos    []primitive.ObjectID `json:"addPhotos"`
	RemovePhotos []primitive.ObjectID `json:"removePhotos"`
}
