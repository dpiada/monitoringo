package main

import (

	"monitoringo/src/routes"
	"github.com/gin-gonic/gin"
) 

func main() {
	server := gin.Default()
    server.GET("/health", routes.HealthHandler)

    server.Run()
}