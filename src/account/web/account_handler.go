package web

import (
	"net/http"
	"strconv"

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

// CreateAccount is a Gin handler that parses the JSON request to create a new account.
// It calls the service layer and returns the new account number as JSON.
//
// Method: POST /accounts
func (ah *accountHandler) CreateAccount(c *gin.Context) {
	// parse JSON etc.

	var req model.CreateAccountReq

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "invalid request"},
		)
		return
	}

	accounNumber, err := ah.svc.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "CreateAccount Success!", "account_number": accounNumber})
}

// GetAccount is a Gin handler to fetch account details by account number.
//
// Method: GET /accounts/:id
func (ah *accountHandler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	accountNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid account number"})
		return
	}

	account, err := ah.svc.GetAccount(accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	c.JSON(http.StatusOK, account)
}
