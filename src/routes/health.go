package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HealthHandler(c *gin.Context) {
	client := c.MustGet("mongo_client").(*mongo.Client)

	err := client.Ping(c, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB: ", err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "K0"})
	}
	log.Println("Successfully connected to MongoDB")

	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}
