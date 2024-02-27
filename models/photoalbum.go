package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PhotoAlbum struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	CoverImageUrl string               `json:"cover_image_url" bson:"cover_image_url,omitempty"`
	Name          string               `json:"name" bson:"name,omitempty"`
	Description   string               `json:"description" bson:"description,omitempty"`
	Owner         primitive.ObjectID   `json:"owner" bson:"owner,omitempty"`
	VisibleTo     []primitive.ObjectID `json:"visible_to" bson:"visible_to,omitempty"`
}
