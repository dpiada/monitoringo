package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResults(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "GET RESULTS"})
}
