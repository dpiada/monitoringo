package main

import (
	"context"
	"log"
	"monitoringo/src/activator"
	"monitoringo/src/database"
	"monitoringo/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server starting...")

	server := gin.Default()
	client := database.InitDatabase()

	// Pass the MongoDB client to the middleware
	server.Use(database.MongoMiddleware(client))

	// Add deferred disconnection with proper context
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Health routes
	server.GET("/health", routes.HealthHandler)
	// Configuration health routes
	server.POST("/configuration", routes.SetHealthCfg)
	server.GET("/configurations", routes.GetHealthCfgList)

	// Get health result routes
	server.GET("/health-results", routes.GetResults)

	go activator.Activator(client)

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
