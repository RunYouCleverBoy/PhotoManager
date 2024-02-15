package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client      *mongo.Client
	db          *mongo.Database
	collections *collections
}

type collections struct {
	users  *mongo.Collection
	Photos *mongo.Collection
}

func NewDb(url string, dbName string) (*Database, error) {
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	db := client.Database(dbName)
	if err != nil {
		return nil, err
	}

	return &Database{
		client:      client,
		db:          db,
		collections: initCollections(db),
	}, nil
}

func (d *Database) Close() {
	d.client.Disconnect(context.Background())
}

func initCollections(db *mongo.Database) *collections {
	col := &collections{}
	col.users = db.Collection("Users")
	col.Photos = db.Collection("Photos")

	initUserSchema(col.users)
	initPhotoSchema(col.Photos)

	return col
}
