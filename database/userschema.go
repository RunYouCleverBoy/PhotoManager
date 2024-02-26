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

type UserCollection struct {
	users *mongo.Collection
}

func (d *UserCollection) GetAll() ([]models.User, error) {
	ctx := context.Background()
	cursor, err := d.users.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (d *UserCollection) Get(id *primitive.ObjectID) (*models.User, error) {
	ctx := context.Background()
	return d.getUserById(&ctx, id)
}

func (d *UserCollection) Create(user *models.User) (*models.User, error) {
	ctx := context.Background()
	if count, _ := d.users.EstimatedDocumentCount(ctx); count == 0 {
		r := models.RoleAdmin
		user.Role = &r
	} else {
		r := models.RoleUser
		user.Role = &r
	}

	if user.ID.IsZero() {
		id := primitive.NewObjectID()
		user.ID = id
	}

	insertResult, err := d.users.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}

	objectId := insertResult.InsertedID.(primitive.ObjectID)
	return d.getUserById(&ctx, &objectId)
}

func (d *UserCollection) Update(id *primitive.ObjectID, user *models.User) (*models.User, error) {
	ctx := context.Background()

	var result User = User{}
	filter := bson.D{{Key: "_id", Value: *id}}
	err := d.users.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: &user}}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *UserCollection) UpdateCredentials(id *primitive.ObjectID, password *string, token *string, refreshToken *string) (*models.User, error) {
	ctx := context.Background()

	var result User = User{}
	filter := bson.D{{Key: "_id", Value: id}}
	updates := bson.D{}
	if password != nil {
		updates = append(updates, bson.E{Key: "password", Value: *password})
	}
	if token != nil {
		updates = append(updates, bson.E{Key: "token", Value: *token})
	}
	if refreshToken != nil {
		updates = append(updates, bson.E{Key: "refreshToken", Value: *refreshToken})
	}
	update := bson.D{{Key: "$set", Value: updates}}
	err := d.users.FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (d *UserCollection) Delete(id *primitive.ObjectID) error {
	ctx := context.Background()
	filter := bson.D{{Key: "_id", Value: id}}
	if _, err := d.users.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (d *UserCollection) getUserById(ctx *context.Context, id *primitive.ObjectID) (*models.User, error) {
	user := User{}
	if err := d.users.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *UserCollection) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	user := User{}
	if err := d.users.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user); err != nil {
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

func initUserSchema(collection *UserCollection) {
	collection.users.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}
