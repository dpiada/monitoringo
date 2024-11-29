package main

import (

	"monitoringo/src/routes"
	"github.com/gin-gonic/gin"
) 

func main() {
	server := gin.Default()
	//Health routes
    server.GET("/health", routes.HealthHandler)

	//Configuration health routes
	server.POST("/configuration", routes.SetHealthCfg)
	server.PUT("/configuration/:id", routes.EditHealthCfg)
	server.DELETE("/configuration/:id", routes.DelHealthCfg)
	server.GET("/configuration/:id", routes.GetHealthCfg)

	//get health result routes
	server.GET("/health-results", routes.GetResults)
	

    server.Run()
}