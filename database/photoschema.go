package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playgrounds.com/models"
)

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

type PhotoModel = models.PhotoModel

func (d *Database) CreatePhoto(photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	_, err := d.collections.Photos.InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (d *Database) GetPhotoById(id string) (*models.PhotoModel, error) {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return nil, err
	}
	return d.getPhotoById(&ctx, objID)
}

func (d *Database) UpdatePhoto(id string, photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = d.collections.Photos.ReplaceOne(ctx, filter, photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (d *Database) DeletePhoto(id string) error {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = d.collections.Photos.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) getPhotoById(ctx *context.Context, id primitive.ObjectID) (*models.PhotoModel, error) {
	photo := PhotoModel{}
	err := d.collections.Photos.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}
