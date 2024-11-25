package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}