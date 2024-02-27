package database

import (
	"testing"
	"time"

	"playgrounds.com/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPhotoSearchCriteriaFromOptions(t *testing.T) {
	selfID := primitive.NewObjectID()
	shotAfter := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	shotBefore := time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC).Unix()
	modifiedAround := time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC).Unix()
	camera := "Canon"
	location := PhotoFilterLocation{
		Geolocation: models.Geolocation{Latitude: 40.7128, Longitude: -74.0060},
		Radius:      1000,
	}
	locationContains := "New York"
	commentsContaining := "great photo"
	workflowStage := "processed"
	upvoteGrade := int8(5)
	isPublic := true
	onlyMine := true

	options := &models.PhotoSearchOptions{
		ShotAfter:          &shotAfter,
		ShotBefore:         &shotBefore,
		ModifiedAround:     &modifiedAround,
		Camera:             &camera,
		Location:           &location,
		LocationContains:   &locationContains,
		CommentsContaining: &commentsContaining,
		OwnedPhotoFilter: &OwnedPhotoFilter{
			WorkflowStage: &workflowStage,
			UpvoteGrade:   &upvoteGrade,
			IsPublic:      &isPublic,
			OnlyMine:      &onlyMine,
		},
	}

	expectedBuilder := bsonBuilder{
		data: bson.D{
			{Key: "shot_date", Value: bson.D{{Key: "$gte", Value: shotAfter}}},
			{Key: "shot_date", Value: bson.D{{Key: "$lte", Value: shotBefore}}},
			{Key: "process_date", Value: bson.D{{Key: "$gte", Value: modifiedAround}}},
			{Key: "camera", Value: bson.D{{Key: "$regex", Value: camera}}},
			{Key: "location", Value: bson.D{
				{Key: "$near", Value: bson.D{
					{Key: "$geometry", Value: location.Geolocation},
					{Key: "$maxDistance", Value: location.Radius},
				}},
			}},
			{Key: "$or", Value: []bson.D{
				{{Key: "place.name", Value: bson.D{{Key: "$regex", Value: locationContains}}}},
				{{Key: "place.city", Value: bson.D{{Key: "$regex", Value: locationContains}}}},
				{{Key: "place.country", Value: bson.D{{Key: "$regex", Value: locationContains}}}},
				{{Key: "place.aliases", Value: bson.D{{Key: "$regex", Value: locationContains}}}},
			}},
			{Key: "comments.comment", Value: bson.D{{Key: "$regex", Value: commentsContaining}}},
			{Key: "workflow.stage", Value: workflowStage},
			{Key: "workflow.upvote_grade", Value: upvoteGrade},
			{Key: "is_public", Value: isPublic},
			{Key: "owner", Value: bson.D{{Key: "$eq", Value: selfID}}},
		},
	}

	result := photoSearchCriteriaFromOptions(&selfID, options)
	assert.Equal(t, expectedBuilder.build(), result)
}

type PhotoFilterLocation = struct {
	Geolocation models.Geolocation `json:"geolocation,omitempty"`
	Radius      float64            `json:"radius,omitempty"`
}

type OwnedPhotoFilter = struct {
	OnlyMine      *bool   `json:"only_mine,omitempty"`
	IsPublic      *bool   `json:"is_public,omitempty"`
	UpvoteGrade   *int8   `json:"upvote_grade,omitempty"`
	WorkflowStage *string `json:"workflow_stage,omitempty"`
}
