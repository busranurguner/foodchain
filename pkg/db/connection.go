package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigDB() *mongo.Database {
	ctx := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("URI"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("DATABASE_NAME")
}
