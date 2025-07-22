package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sbdoddi7/innoscripta/src/account/repository"
	"github.com/sbdoddi7/innoscripta/src/account/service"
	"github.com/sbdoddi7/innoscripta/src/account/web"
	"github.com/sbdoddi7/innoscripta/src/config"
	"github.com/sbdoddi7/innoscripta/src/platform/database"
	txweb "github.com/sbdoddi7/innoscripta/src/transaction/web"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	cof := config.LoadConfig()

	db, err := database.NewPostgressDb(cof.PostgresDSN)
	if err != nil {
		log.Printf("Failed to connect postgress database: %v", err)
		return nil
	}
	repo := repository.NewAccountRepository(db)
	svc := service.NewAccountService(repo)
	handler := web.NewAccountHandler(svc)

	r.POST("/accounts", handler.CreateAccount)
	r.GET("/accounts/:id", handler.GetAccount)
	r.POST("/transactions", txweb.CreateTransaction)
	r.GET("/accounts/:id/transactions", txweb.GetTransactions)

	return r
}
