package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SetHealthCfg(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "SET"})
}

func EditHealthCfg(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "UPDATE on "+string(id)})
}

func GetHealthCfg(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "GET on "+string(id)})
}

func DelHealthCfg(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "DEL on "+string(id)})
}