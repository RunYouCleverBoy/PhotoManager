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

type UserDocument struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

var ErrNoDocuments = mongo.ErrNoDocuments
var ErrInvalidId = errors.New("invalid id")
var ErrInvalidUser = errors.New("invalid user")

func FromUserNoID(user *models.User) UserDocument {
	return UserDocument{
		Name:  user.Name,
		Email: user.Email,
	}
}

func FromUser(user *models.User) UserDocument {
	id, _ := objectIdFromHex(user.ID)
	return UserDocument{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (u UserDocument) ToUser() models.User {
	return models.User{
		ID:    u.ID.Hex(),
		Name:  u.Name,
		Email: u.Email,
	}
}

func (d *Database) GetAll() ([]models.User, error) {
	ctx := context.Background()
	cursor, err := d.collections.users.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	for cursor.Next(ctx) {
		var user UserDocument
		err := cursor.Decode(&user)
		if err != nil {
			continue
		}
		users = append(users, user.ToUser())
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
	insertedId, err := d.collections.users.InsertOne(ctx, FromUserNoID(&user))
	if err != nil {
		return nil, err
	}

	id := insertedId.InsertedID.(primitive.ObjectID)
	return d.getUserById(&ctx, id)
}

func (d *Database) Update(id string, user models.User) (*models.User, error) {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidId
	}

	userDocument := UserDocument{}
	filter := bson.D{{Key: "_id", Value: objID}}
	err = d.collections.users.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: FromUser(&user)}}).Decode(&userDocument)
	if err != nil {
		return nil, err
	}

	updatedUser := userDocument.ToUser()
	return &updatedUser, nil
}

func (d *Database) Delete(id string) error {
	ctx := context.Background()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = d.collections.users.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) getUserById(ctx *context.Context, id primitive.ObjectID) (*models.User, error) {
	userDocument := UserDocument{}
	err := d.collections.users.FindOne(*ctx, bson.D{{Key: "_id", Value: id}}).Decode(&userDocument)
	if err != nil {
		return nil, err
	} else {
		user := userDocument.ToUser()
		return &user, nil
	}
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
