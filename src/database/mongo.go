package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Client

func ConnectDatabase() {
	uri := os.Getenv("MONGODB_URI")

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {

		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	mongodb = client
}
