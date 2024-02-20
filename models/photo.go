package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkFlow struct {
	UpvoteGrade   int    `json:"upvote_grade" bson:"upvote_grade,omitempty"`
	WorkflowStage string `json:"workflow_stage" bson:"workflow_stage,omitempty"`
}

type PhotoMetadata struct {
	ShotDate     int64       `json:"shot_date" bson:"shot_date"`
	ModifiedDate int64       `json:"process_date" bson:"process_date"`
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
	Time          int64              `json:"time" bson:"time"`
}

type PhotoModel struct {
	ID        primitive.ObjectID   `json:"id" bson:"id"`
	Url       string               `json:"url" bson:"url"`
	IsPublic  bool                 `json:"is_public" bson:"is_public"`
	Owner     primitive.ObjectID   `json:"owner" bson:"owner"`
	VisibleTo []primitive.ObjectID `json:"visible_to" bson:"visible_to"`
	Metadata  PhotoMetadata        `json:"metadata" bson:"metadata"`
	WorkFlow  WorkFlow             `json:"workflow" bson:"workflow_stage"`
	SimilarTo []primitive.ObjectID `json:"similar_to" bson:"similar_to"`
	Ancestor  primitive.ObjectID   `json:"ancestor" bson:"ancestor"`
	Comments  []Comments           `json:"comments" bson:"comments"`
}

type PhotoSearchOptions struct {
	ShotAfter      *int64  `json:"shot_after,omitempty"`
	ShotBefore     *int64  `json:"shot_before,omitempty"`
	ModifiedAround *int64  `json:"modified_around,omitempty"`
	Camera         *string `json:"camera,omitempty"`
	Location       *struct {
		Geolocation Geolocation `json:"geolocation,omitempty"`
		Radius      float64     `json:"radius,omitempty"`
	} `json:"location,omitempty"`
	LocationContains   *string `json:"location_contains,omitempty"`
	CommentsContaining *string `json:"comments_containing,omitempty"`
	OwnedPhotoFilter   *struct {
		OnlyMine      *bool   `json:"only_mine,omitempty"`
		IsPublic      *bool   `json:"is_public,omitempty"`
		UpvoteGrade   *int8   `json:"upvote_grade,omitempty"`
		WorkflowStage *string `json:"workflow_stage,omitempty"`
	} `json:"owned_photo_filter,omitempty"`
}
