package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playgrounds.com/models"
)

func initPhotoSchema(collection *PhotosCollection) {
	photos := collection.Photos
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

type PhotosCollection struct {
	Photos *mongo.Collection
}

func (d *PhotosCollection) CreatePhoto(photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	_, err := d.Photos.InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (d *PhotosCollection) GetPhotos(callerObjectId *primitive.ObjectID, searchCriteria *models.PhotoSearchOptions) ([]models.PhotoModel, error) {
	visibilityCriteria := bson.D{{Key: "$or", Value: []bson.D{
		{{Key: "is_public", Value: true}},
		{{Key: "visible_to", Value: callerObjectId}},
		{{Key: "owner", Value: callerObjectId}},
	}}}

	filter := bson.D{{Key: "$and", Value: []bson.D{visibilityCriteria, *photoSearchCriteriaFromOptions(callerObjectId, searchCriteria)}}}
	ctx := context.Background()
	cursor, err := d.Photos.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	photos := make([]models.PhotoModel, 0)
	for cursor.Next(ctx) {
		var photo models.PhotoModel
		err = cursor.Decode(&photo)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (d *PhotosCollection) GetPhotoById(id string) (*models.PhotoModel, error) {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return nil, err
	}
	return d.getPhotoById(&ctx, objID)
}

func (d *PhotosCollection) UpdatePhoto(id string, photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = d.Photos.ReplaceOne(ctx, filter, photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (d *PhotosCollection) DeletePhoto(id string) error {
	ctx := context.Background()
	objID, err := objectIdFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = d.Photos.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (d *PhotosCollection) getPhotoById(ctx *context.Context, id primitive.ObjectID) (*models.PhotoModel, error) {
	photo := PhotoModel{}
	err := d.Photos.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func photoSearchCriteriaFromOptions(selfId *primitive.ObjectID, options *models.PhotoSearchOptions) *bson.D {
	bsonBuilder := bsonBuilder{data: bson.D{}}

	bsonBuilder.addValIf("shot_date", options.ShotAfter != nil, bson.D{{Key: "$gte", Value: *options.ShotAfter}})
	bsonBuilder.addValIf("shot_date", options.ShotBefore != nil, bson.D{{Key: "$lte", Value: *options.ShotBefore}})
	bsonBuilder.addValIf("process_date", options.ModifiedAround != nil, bson.D{{Key: "$gte", Value: *options.ModifiedAround}})
	bsonBuilder.addIfContains("camera", options.Camera)

	if location := options.Location; location != nil {
		bsonBuilder.addVal("location", bson.D{{Key: "$near", Value: bson.D{
			{Key: "$geometry", Value: location.Geolocation},
			{Key: "$maxDistance", Value: location.Radius},
		}}})
	}
	if pLocation := options.LocationContains; pLocation != nil {
		locationStr := *pLocation
		bsonBuilder.addOr(
			bson.D{{Key: "place.name", Value: bson.D{{Key: "$regex", Value: locationStr}}}},
			bson.D{{Key: "place.city", Value: bson.D{{Key: "$regex", Value: locationStr}}}},
			bson.D{{Key: "place.country", Value: bson.D{{Key: "$regex", Value: locationStr}}}},
			bson.D{{Key: "place.aliases", Value: bson.D{{Key: "$regex", Value: locationStr}}}},
		)
	}

	bsonBuilder.addIfContains("comments.comment", options.CommentsContaining)

	if pOpts := options.OwnedPhotoFilter; pOpts != nil {
		ownerOpts := *pOpts
		bsonBuilder.addValIf("workflow.stage", ownerOpts.WorkflowStage != nil, *ownerOpts.WorkflowStage)
		bsonBuilder.addValIf("workflow.upvote_grade", ownerOpts.UpvoteGrade != nil, *ownerOpts.UpvoteGrade)
		bsonBuilder.addValIf("is_public", ownerOpts.IsPublic != nil, *ownerOpts.IsPublic)
		if options.OwnedPhotoFilter.OnlyMine != nil && *options.OwnedPhotoFilter.OnlyMine {
			bsonBuilder.addVal("owner", bson.D{{Key: "$eq", Value: *selfId}})
		}
	}

	return bsonBuilder.build()
}
