package photoalbums

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddOrRemoveVisibilityRequestBody struct {
	AddVisibleTo    []primitive.ObjectID `json:"addVisibleTo"`
	RemoveVisibleTo []primitive.ObjectID `json:"removeVisibleTo"`
}
