package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/innoscripta/src/account/web"
	acweb "github.com/sbdoddi7/innoscripta/src/account/web"
	txweb "github.com/sbdoddi7/innoscripta/src/transaction/web"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	repo := 
	svc :=
	handler := web.NewAccountHandler(svc)

	r.POST("/accounts", acweb.CreateAccount)
	r.GET("/accounts/:id", acweb.GetAccount)
	r.POST("/transactions", txweb.CreateTransaction)
	r.GET("/accounts/:id/transactions", txweb.GetTransactions)

	return r
}
