package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"playgrounds.com/models"
)

var ErrNoDocuments = mongo.ErrNoDocuments
var ErrInvalidId = errors.New("invalid id")
var ErrInvalidUser = errors.New("invalid user")

type User = models.User

func (d *Database) GetAll() ([]models.User, error) {
	ctx := context.Background()
	cursor, err := d.collections.users.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (d *Database) Get(id string) (*models.User, error) {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}
	return d.getUserById(&ctx, objID)
}

func (d *Database) Create(user models.User) (*models.User, error) {
	ctx := context.Background()
	insertResult, err := d.collections.users.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}

	id := insertResult.InsertedID.(primitive.ObjectID)
	return d.getUserById(&ctx, id)
}

func (d *Database) Update(id string, user models.User) (*models.User, error) {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	var result User = User{}
	filter := bson.D{{Key: "_id", Value: objID}}
	err = d.collections.users.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: &user}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *Database) Delete(id string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	if _, err = d.collections.users.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (d *Database) getUserById(ctx *context.Context, id primitive.ObjectID) (*models.User, error) {
	user := User{}
	if err := d.collections.users.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func objectIdFromHex(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, ErrInvalidId
	}
	return objID, nil
}

func initUserSchema(collection *mongo.Collection) {
	collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
