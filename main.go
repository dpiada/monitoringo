package main

import (
	"context"
	"log"
	"monitoringo/src/database"
	"monitoringo/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	client := database.InitDatabase()
	defer client.Disconnect(context.Background())

	// Health routes
	server.GET("/health", routes.HealthHandler)

	// Configuration health routes
	server.POST("/configuration", routes.SetHealthCfg)
	server.PUT("/configuration/:id", routes.EditHealthCfg)
	server.DELETE("/configuration/:id", routes.DelHealthCfg)
	server.GET("/configuration/:id", routes.GetHealthCfg)

	// Get health result routes
	server.GET("/health-results", routes.GetResults)

	server.Run(":8080")
	log.Println("Server starting...")
}
