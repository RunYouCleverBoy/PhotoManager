package database

import (
	"context"
	"log"
	"sync"

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
	log.Printf("Database: Connecting to database at %s", url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	log.Print("Database: Connected to database")

	db := client.Database(dbName)
	database = &databaseContext{
		client: client,
	}

	log.Print("Database: Database initialized")

	collections := initCollections(db)
	return &Database{
		db:          db,
		collections: collections,
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

	waitGroup := sync.WaitGroup{}

	type initSchemaJob struct {
		name string
		job  func() error
	}

	jobs := []initSchemaJob{
		{"Users", func() error { return initUserSchema(col.users) }},
		{"Photos", func() error { return initPhotoSchema(col.photos) }},
		{"Albums", func() error { return initAlbumSchema(col.albums) }},
	}

	for _, job := range jobs {
		waitGroup.Add(1)
		go func(job initSchemaJob) {
			if err := job.job(); err == nil {
				log.Printf("Database: Schema for %s initialized", job.name)
			} else {
				log.Panicf("Database: Schema for %s failed\n%s", job.name, err)
			}
			waitGroup.Done()
		}(job)
	}

	waitGroup.Wait()

	return col
}
