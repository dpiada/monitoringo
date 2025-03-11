package routes

import (
	"context"
	"net/http"
	"time"

	"monitoringo/src/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetHealthCfg creates a new health configuration
func SetHealthCfg(c *gin.Context) {

	var body models.HealthCfgRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	client := c.MustGet("mongo_client").(*mongo.Client)

	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	result, err := client.Database("configurations").Collection("health_config").InsertOne(ctx, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Health configuration created successfully",
		"id":      result.InsertedID,
	})
}

// GetHealthCfgList retrieves a list of all health configurations
func GetHealthCfgList(c *gin.Context) {
	client := c.MustGet("mongo_client").(*mongo.Client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := client.Database("configurations").Collection("health_config").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	var healthConfigs []models.HealthCfgResponse
	if err := cursor.All(ctx, &healthConfigs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Health configurations retrieved successfully",
		"configurations": healthConfigs,
		"count":          len(healthConfigs),
	})
}
