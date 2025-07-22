package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	// parse JSON etc.
	c.JSON(http.StatusOK, gin.H{"message": "CreateTransaction not implemented"})
}

func GetTransactions(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"account_id": id, "message": "GetTransactions not implemented"})
}
