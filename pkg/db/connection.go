package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

var DB *mongo.Client = MongoConnection()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Example").Collection(collectionName)
	return collection
}
