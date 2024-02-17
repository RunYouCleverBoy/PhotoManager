package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkFlow struct {
	IsRejected    bool   `json:"is_rejected" bson:"is_rejected"`
	IsAccepted    bool   `json:"is_accepted" bson:"is_accepted"`
	WorkflowStage string `json:"workflow_stage" bson:"workflow_stage"`
}

type PhotoMetadata struct {
	ShotDate     time.Time   `json:"shot_date" bson:"shot_date"`
	ModifiedDate time.Time   `json:"process_date" bson:"process_date"`
	Camera       string      `json:"camera" bson:"camera"`
	Location     Geolocation `json:"location" bson:"location"`
	Place        Place       `json:"place" bson:"place"`
	Exposure     string      `json:"exposure" bson:"exposure"`
	FNumber      float32     `json:"f_number" bson:"f_number"`
	ISO          int         `json:"iso" bson:"iso"`
	Description  string      `json:"description" bson:"description"`
}

type Comments struct {
	CommenterID   primitive.ObjectID `json:"commenter_id" bson:"commenter_id"`
	CommenterName string             `json:"commenter_name" bson:"commenter_name"`
	Comment       string             `json:"comment" bson:"comment"`
	Time          time.Time          `json:"time" bson:"time"`
}

type PhotoModel struct {
	ID        string               `json:"id" bson:"id"`
	Url       string               `json:"url" bson:"url"`
	IsPublic  bool                 `json:"is_public" bson:"is_public"`
	Owner     primitive.ObjectID   `json:"owner" bson:"owner"`
	VisibleTo []primitive.ObjectID `json:"visible_to" bson:"visible_to"`
	Metadata  PhotoMetadata        `json:"metadata" bson:"metadata"`
	WorkFlow  WorkFlow             `json:"workflow_stage" bson:"workflow_stage"`
	SimilarTo []primitive.ObjectID `json:"similar_to" bson:"similar_to"`
	Ancestor  primitive.ObjectID   `json:"ancestor" bson:"ancestor"`
	Comments  []Comments           `json:"comments" bson:"comments"`
}
