package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/innoscripta/src/model"
)

type accountHandler struct {
	svc model.AccountService
}

func NewAccountHandler(svc model.AccountService) *accountHandler {
	return &accountHandler{
		svc: svc,
	}
}

func (ah *accountHandler) CreateAccount(c *gin.Context) {
	// parse JSON etc.
	err := ah.svc.CreateAccount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "CreateAccount not implemented"})
}

func (ah *accountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"account_id": id, "message": "GetAccount not implemented"})
}
