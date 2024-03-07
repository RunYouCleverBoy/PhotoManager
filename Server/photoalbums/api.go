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

type AlbumSearchCriteria struct {
	OwnerID          *primitive.ObjectID `json:"owner_id,omitempty"`
	NameRegex        *string             `json:"name,omitempty"`
	DescriptionRegex *string             `json:"description,omitempty"`
	VisibilityTo     *primitive.ObjectID `json:"visibility_to,omitempty"`
}
