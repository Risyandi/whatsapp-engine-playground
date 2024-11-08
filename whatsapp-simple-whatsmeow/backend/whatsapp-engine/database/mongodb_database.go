package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	URL          string
	DatabaseName string
	Client       *mongo.Client
}

func (mongoDb *MongoDB) Connect() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	newClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDb.URL))
	if err != nil {
		log.Fatal(err)
	}

	mongoDb.Client = newClient
	return newClient
}

func (mongoDb *MongoDB) Database(databaseName string) *mongo.Database {
	return mongoDb.Client.Database(databaseName)
}

func (mongoDb *MongoDB) Collection(collection string) *mongo.Collection {
	return mongoDb.Client.Database(mongoDb.DatabaseName).Collection(collection)
}
