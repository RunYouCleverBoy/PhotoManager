package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playgrounds.com/models"
)

type photoMetadata struct {
	ShotDate time.Time    `bson:"shot_date"`
	Camera   string       `bson:"camera"`
	Location GeoJson      `bson:"location"`
	Place    models.Place `bson:"place"`
	Exposure string       `bson:"exposure"`
	FNumber  float32      `bson:"f_number"`
	ISO      int          `bson:"iso"`
}

type PhotoDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Url         string             `bson:"url"`
	Description string             `bson:"description"`
	IsPublic    bool               `bson:"is_public"`
	Metadata    photoMetadata      `bson:"metadata"`
	WorkFlow    models.WorkFlow    `bson:"workflow_stage"`
}

func PhotoDocumentFromModel(p models.PhotoModel) PhotoDocument {
	return PhotoDocument{
		Url:         p.Url,
		Description: p.Description,
		IsPublic:    p.IsPublic,
		WorkFlow:    p.WorkFlow,
		Metadata: photoMetadata{
			ShotDate: p.Metadata.ShotDate,
			Camera:   p.Metadata.Camera,
			Location: fromGeolocation(p.Metadata.Location),
			Place:    p.Metadata.Place,
			Exposure: p.Metadata.Exposure,
			FNumber:  p.Metadata.FNumber,
			ISO:      p.Metadata.ISO,
		},
	}
}

func (p PhotoDocument) ToModel() models.PhotoModel {
	return models.PhotoModel{
		ID:          p.ID.Hex(),
		Url:         p.Url,
		Description: p.Description,
		IsPublic:    p.IsPublic,
		WorkFlow:    p.WorkFlow,
		Metadata: models.PhotoMetadata{
			ShotDate: p.Metadata.ShotDate,
			Camera:   p.Metadata.Camera,
			Location: p.Metadata.Location.toGeolocation(),
			Exposure: p.Metadata.Exposure,
			FNumber:  p.Metadata.FNumber,
			ISO:      p.Metadata.ISO,
		},
	}
}

func initPhotoSchema(photos *mongo.Collection) {
	ctx := context.Background()
	photos.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "url", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	photos.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "workflow_stage", Value: 1}},
	})
}
