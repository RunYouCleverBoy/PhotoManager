package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *databaseContext

type databaseContext struct {
	client *mongo.Client
}

type Database struct {
	db          *mongo.Database
	collections *Collections
}

func (d *Database) UserCollection() *UserCollection {
	return d.collections.users
}
func (d *Database) PhotosCollection() *PhotosCollection {
	return d.collections.photos
}

type Collections struct {
	users  *UserCollection
	photos *PhotosCollection
}

func NewDb(url string, dbName string) (*Database, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	db := client.Database(dbName)
	database = &databaseContext{
		client: client,
	}
	if err != nil {
		return nil, err
	}

	return &Database{
		db:          db,
		collections: initCollections(db),
	}, nil
}

func (d *Database) Close() {
	database.client.Disconnect(context.Background())
}

func initCollections(db *mongo.Database) *Collections {
	col := &Collections{}
	col.users = &UserCollection{db.Collection("Users")}
	col.photos = &PhotosCollection{db.Collection("Photos")}

	initUserSchema(col.users)
	initPhotoSchema(col.photos)

	return col
}
