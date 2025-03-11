package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variable should be capitalized since it's used in exported functions
var mongoClient *mongo.Client

func InitDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Declaring 'client' here shadows the global variable and causes it to remain nil
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}

	// Ping the database to ensure connectivity
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
	}

	// Set the global variable
	mongoClient = client
	return client
}

func GetClient() *mongo.Client {
	return mongoClient
}

// Update the middleware to accept a client parameter
func MongoMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mongo_client", client)
		c.Next()
	}
}
