package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	WorkflowStageFootage    = "footage"
	WorkflowStageCollection = "collection"
	WorkflowStageAlbum      = "album"
)

type WorkFlow struct {
	UpvoteGrade   int                  `json:"upvote_grade" bson:"upvote_grade,omitempty"`
	WorkflowStage string               `json:"workflow_stage" bson:"workflow_stage,omitempty"`
	Albums        []primitive.ObjectID `json:"albums" bson:"albums,omitempty"`
}

type PhotoMetadata struct {
	ShotDate     *int64       `json:"shot_date,omitempty" bson:"shot_date,omitempty"`
	ModifiedDate *int64       `json:"process_date,omitempty" bson:"process_date,omitempty"`
	Camera       *string      `json:"camera,omitempty" bson:"camera,omitempty"`
	Location     *Geolocation `json:"location,omitempty" bson:"location,omitempty"`
	Place        *Place       `json:"place,omitempty" bson:"place,omitempty"`
	Exposure     *string      `json:"exposure,omitempty" bson:"exposure,omitempty"`
	FNumber      *float32     `json:"f_number,omitempty" bson:"f_number,omitempty"`
	ISO          *int         `json:"iso,omitempty" bson:"iso,omitempty"`
	Description  *string      `json:"description,omitempty" bson:"description,omitempty"`
}

type Comments struct {
	CommenterID   primitive.ObjectID `json:"commenter_id" bson:"commenter_id"`
	CommenterName string             `json:"commenter_name" bson:"commenter_name"`
	Comment       string             `json:"comment" bson:"comment"`
	Time          int64              `json:"time" bson:"time"`
}

type PhotoModel struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Url       string               `json:"url" bson:"url,omitempty"`
	IsPublic  bool                 `json:"is_public" bson:"is_public,omitempty"`
	Owner     primitive.ObjectID   `json:"owner" bson:"owner,omitempty"`
	VisibleTo []primitive.ObjectID `json:"visible_to" bson:"visible_to,omitempty"`
	Metadata  PhotoMetadata        `json:"metadata" bson:"metadata,omitempty"`
	WorkFlow  WorkFlow             `json:"workflow" bson:"workflow_stage,omitempty"`
	SimilarTo []primitive.ObjectID `json:"similar_to" bson:"similar_to,omitempty"`
	Ancestor  primitive.ObjectID   `json:"ancestor" bson:"ancestor,omitempty"`
	Comments  []Comments           `json:"comments" bson:"comments,omitempty"`
	Tags      []string             `json:"tags" bson:"tags,omitempty"`
}

type PhotoSearchLocation struct {
	Geolocation Geolocation `json:"geolocation,omitempty"`
	Radius      float64     `json:"radius,omitempty"`
}

type UpvoteGradeRange struct {
	Min int8 `json:"min,omitempty"`
	Max int8 `json:"max,omitempty"`
}

type PhotoSearchOwnedPhotoFilter struct {
	OnlyMine      *bool             `json:"only_mine,omitempty"`
	IsPublic      *bool             `json:"is_public,omitempty"`
	UpvoteGrade   *UpvoteGradeRange `json:"upvote_grade_min,omitempty"`
	WorkflowStage *string           `json:"workflow_stage,omitempty"`
}

type PhotoSearchOptions struct {
	ShotAfter          *int64                       `json:"shot_after,omitempty"`
	ShotBefore         *int64                       `json:"shot_before,omitempty"`
	ModifiedAround     *int64                       `json:"modified_around,omitempty"`
	Camera             *string                      `json:"camera,omitempty"`
	Location           *PhotoSearchLocation         `json:"location,omitempty"`
	LocationContains   *string                      `json:"location_contains,omitempty"`
	CommentsContaining *string                      `json:"comments_containing,omitempty"`
	OwnedPhotoFilter   *PhotoSearchOwnedPhotoFilter `json:"owned_photo_filter,omitempty"`
}
