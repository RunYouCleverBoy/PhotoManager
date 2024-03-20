package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playgrounds.com/models"
	"playgrounds.com/utils"
)

var (
	ErrInvisibleToUser = errors.New("photo is not visible to user")
)

func initPhotoSchema(collection *PhotosCollection) error {
	photos := collection.Photos
	ctx := context.Background()
	if _, err := photos.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "url", Value: 1}},
		Options: options.Index().SetUnique(true),
	}); err != nil {
		return err
	}
	if _, err := photos.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "workflow_stage", Value: 1}},
	}); err != nil {
		return err
	}
	return nil
}

type PhotoModel = models.PhotoModel

type PhotosCollection struct {
	Photos *mongo.Collection
}

func (d *PhotosCollection) CreatePhoto(ownerId *primitive.ObjectID, photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	_, err := d.Photos.InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func (d *PhotosCollection) SetPhotoFile(id *primitive.ObjectID, fileUrl string) error {
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: *id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "url", Value: fileUrl}}}}
	_, err := d.Photos.UpdateOne(ctx, filter, update)
	return err
}

func (d *PhotosCollection) AddCommentToPhoto(photoId *primitive.ObjectID, comment *models.Comments) (*models.Comments, error) {
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: *photoId}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "comments", Value: *comment}}}}
	_, err := d.Photos.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (d *PhotosCollection) AddTagToPhoto(photoId *primitive.ObjectID, tag string) (*string, error) {
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: *photoId}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "tags", Value: tag}}}}
	_, err := d.Photos.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (d *PhotosCollection) GetPhotosByUserId(userId *primitive.ObjectID, indexRange *utils.IntRange[int]) ([]models.PhotoModel, error) {
	ctx := context.Background()
	options := options.Find()
	if !indexRange.IsNullOrEmpty() {
		options = options.SetSkip(int64(indexRange.Start())).SetLimit(int64(indexRange.Length()))
	}

	cursor, err := d.Photos.Find(ctx, bson.D{{Key: "owner", Value: *userId}}, options)
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

func (d *PhotosCollection) GetPhotos(callerObjectId *primitive.ObjectID, searchCriteria *models.PhotoSearchOptions, indexRange *utils.IntRange[int]) ([]models.PhotoModel, error) {
	visibilityCriteria := bson.D{{Key: "$or", Value: []bson.D{
		{{Key: "is_public", Value: true}},
		{{Key: "visible_to", Value: callerObjectId}},
		{{Key: "owner", Value: callerObjectId}},
	}}}

	filter := bson.D{{Key: "$and", Value: []bson.D{visibilityCriteria, *photoSearchCriteriaFromOptions(callerObjectId, searchCriteria)}}}
	ctx := context.Background()
	options := options.Find()
	if !indexRange.IsNullOrEmpty() {
		options = options.SetSkip(int64(indexRange.Start())).SetLimit(int64(indexRange.Length()))
	}

	cursor, err := d.Photos.Find(ctx, filter, options)
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

func (d *PhotosCollection) GetPhotoById(id *primitive.ObjectID) (*models.PhotoModel, error) {
	ctx := context.Background()
	return d.getPhotoById(&ctx, id)
}

func (d *PhotosCollection) UpdatePhoto(id *primitive.ObjectID, photo models.PhotoModel) (*models.PhotoModel, error) {
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: *id}}
	_, err := d.Photos.ReplaceOne(ctx, filter, photo)
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

func (d *PhotosCollection) AddOrRemoveAlbumFromManyPhotos(albumId *primitive.ObjectID, addPhotos *[]primitive.ObjectID, removePhotos *[]primitive.ObjectID) error {
	photosToAdd := []primitive.ObjectID{}
	photosToRemove := []primitive.ObjectID{}
	if addPhotos != nil {
		photosToAdd = *addPhotos
	}
	if removePhotos != nil {
		photosToRemove = *removePhotos
	}

	addPhotosFilter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: photosToAdd}}}}
	removePhotosFilter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: photosToRemove}}}}
	if _, err := d.Photos.UpdateMany(context.Background(), addPhotosFilter, bson.D{{Key: "$addToSet", Value: bson.D{{Key: "albums", Value: *albumId}}}}); err != nil {
		return err
	}

	if _, err := d.Photos.UpdateMany(context.Background(), removePhotosFilter, bson.D{{Key: "$pull", Value: bson.D{{Key: "albums", Value: *albumId}}}}); err != nil {
		return err
	}

	return nil
}

func (d *PhotosCollection) VerifyVisibilityForAllPhotos(photoIds *[]primitive.ObjectID, userId *primitive.ObjectID) error {
	if photoIds == nil || len(*photoIds) == 0 {
		return nil
	}

	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: *photoIds}}},
		{Key: "$or", Value: []bson.D{
			{{Key: "visible_to", Value: *userId}},
			{{Key: "owner", Value: *userId}},
		}},
	}
	count, err := d.Photos.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if count < int64(len(*photoIds)) {
		return ErrInvisibleToUser
	}
	return nil
}

func (d *PhotosCollection) getPhotoById(ctx *context.Context, id *primitive.ObjectID) (*models.PhotoModel, error) {
	photo := PhotoModel{}
	err := d.Photos.FindOne(*ctx, bson.D{{Key: "_id", Value: *id}}).Decode(&photo)
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
		bsonBuilder.addRangeIf("workflow.upvote_grade", ownerOpts.UpvoteGrade != nil, upvoteRangeConverter(ownerOpts.UpvoteGrade))
		bsonBuilder.addValIf("is_public", ownerOpts.IsPublic != nil, *ownerOpts.IsPublic)
		if ownerOpts.OnlyMine != nil && *ownerOpts.OnlyMine {
			bsonBuilder.addVal("owner", bson.D{{Key: "$eq", Value: *selfId}})
		}
	}

	return bsonBuilder.build()
}

func upvoteRangeConverter(upvoteRange *models.UpvoteGradeRange) *utils.IntRange[int] {
	if upvoteRange == nil {
		return nil
	}
	return &utils.IntRange[int]{Min: int(upvoteRange.Min), Max: int(upvoteRange.Max)}
}
