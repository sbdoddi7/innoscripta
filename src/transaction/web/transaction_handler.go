package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/innoscripta/src/model"
	logger "github.com/sbdoddi7/innoscripta/src/platform/log"
)

type transactionHandler struct {
	service model.TransactionService
}

func NewTransactionHandler(s model.TransactionService) *transactionHandler {
	return &transactionHandler{service: s}
}

func (th *transactionHandler) CreateTransaction(c *gin.Context) {
	var req model.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.WithError(err).Info("bad request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	txID, err := th.service.CreateTransaction(req.AccountNumber, req.Amount, req.Type)
	if err != nil {
		logger.Logger.WithError(err).Info("internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue transaction"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"transaction_id": txID, "status": "queued"})
}

func (th *transactionHandler) GetTransactions(c *gin.Context) {
	accountNumberStr := c.Param("account_number")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	accountNumber, err := strconv.ParseInt(accountNumberStr, 10, 64)
	if err != nil {
		logger.Logger.WithError(err).Warn("invalid customer number")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account_number"})
		return
	}
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	transactions, err := th.service.GetTransactions(accountNumber, int64(page), int64(limit))
	if err != nil {
		logger.Logger.WithError(err).Info("internal server error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
