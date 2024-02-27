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

func (d *Database) AlbumsCollection() *AlbumCollection {
	return d.collections.albums
}

type Collections struct {
	users  *UserCollection
	photos *PhotosCollection
	albums *AlbumCollection
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
	col.albums = &AlbumCollection{db.Collection("Albums")}

	initUserSchema(col.users)
	initPhotoSchema(col.photos)
	initAlbumSchema(col.albums)

	return col
}
