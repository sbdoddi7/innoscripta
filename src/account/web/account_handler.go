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

	var req model.CreateAccountReq

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "invalid request"},
		)
	}

	accounNumber, err := ah.svc.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "CreateAccount Success!", "account_number": accounNumber})
}

func (ah *accountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")

	account, err := ah.svc.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.JSON(http.StatusOK, account)
}
