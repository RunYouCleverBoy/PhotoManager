package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"playgrounds.com/models"
	"playgrounds.com/utils"
)

type AlbumCollection struct {
	Albums *mongo.Collection
}

func initAlbumSchema(collection *AlbumCollection) {
	album := collection.Albums
	ctx := context.Background()
	album.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: 1}},
	})
	album.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "owner", Value: 1}},
	})
}

func (a *AlbumCollection) GetAlbumsBy(criteria *models.AlbumSearchCriteria) (*[]models.PhotoAlbum, error) {
	result := make([]models.PhotoAlbum, 0)

	ctx := context.Background()
	bson := albumSearchCriteriaFromOptions(criteria)
	cursor, err := a.Albums.Find(ctx, bson)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var album models.PhotoAlbum
		if err := cursor.Decode(&album); err != nil {
			return nil, err
		}
		result = append(result, album)
	}

	return &result, nil
}

func (a *AlbumCollection) CreateAlbum(album *models.PhotoAlbum) error {
	ctx := context.Background()
	_, err := a.Albums.InsertOne(ctx, album)
	return err
}

func (a *AlbumCollection) GetAlbumById(id *string) (*models.PhotoAlbum, error) {
	ctx := context.Background()
	var album models.PhotoAlbum
	err := a.Albums.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&album)
	return &album, err
}

func (a *AlbumCollection) AddOrRemoveAlbumVisibility(albumId *primitive.ObjectID, addVisible *[]primitive.ObjectID, removeVisible *[]primitive.ObjectID) error {
	ctx := context.Background()
	if !utils.IsNullOrEmpty(addVisible) {
		_, err := a.Albums.UpdateOne(ctx, bson.D{{Key: "_id", Value: albumId}}, bson.D{{Key: "$addToSet", Value: bson.D{{Key: "visible_to", Value: bson.D{{Key: "$each", Value: addVisible}}}}}})
		if err != nil {
			return err
		}
	}

	if !utils.IsNullOrEmpty(removeVisible) {
		_, err := a.Albums.UpdateOne(ctx, bson.D{{Key: "_id", Value: albumId}}, bson.D{{Key: "$pull", Value: bson.D{{Key: "visible_to", Value: bson.D{{Key: "$in", Value: removeVisible}}}}}})
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AlbumCollection) DeleteAlbum(albumId *primitive.ObjectID) error {
	ctx := context.Background()
	_, err := a.Albums.DeleteOne(ctx, bson.D{{Key: "_id", Value: albumId}})
	return err
}

func albumSearchCriteriaFromOptions(options *models.AlbumSearchCriteria) *bson.D {
	bsonBuilder := bsonBuilder{data: bson.D{}}

	bsonBuilder.addValIf("owner", options.OwnerID != nil, options.OwnerID)
	bsonBuilder.addIfContains("name", options.NameRegex)
	bsonBuilder.addIfContains("description", options.DescriptionRegex)
	bsonBuilder.addValIf("visible_to", options.VisibilityTo != nil, options.VisibilityTo)

	return bsonBuilder.build()
}
