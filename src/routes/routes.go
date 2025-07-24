package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	accRepo "github.com/sbdoddi7/innoscripta/src/account/repository"
	accSvc "github.com/sbdoddi7/innoscripta/src/account/service"
	accWeb "github.com/sbdoddi7/innoscripta/src/account/web"
	"github.com/sbdoddi7/innoscripta/src/platform/queue"
	txRepo "github.com/sbdoddi7/innoscripta/src/transaction/repository"
	txSvc "github.com/sbdoddi7/innoscripta/src/transaction/service"
	txWeb "github.com/sbdoddi7/innoscripta/src/transaction/web"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(postgressDb *sql.DB, mongoClient *mongo.Client, rabbitCh *amqp.Channel) *gin.Engine {
	r := gin.Default()

	queueName := "transactions"
	rabbitCh.QueueDeclare(queueName, true, false, false, false, nil)

	accountRepo := accRepo.NewAccountRepository(postgressDb)
	accountSvc := accSvc.NewAccountService(accountRepo)
	accountWeb := accWeb.NewAccountHandler(accountSvc)

	transactionRepo := txRepo.NewTransactionRepository(mongoClient, "ledger", "transaction_logs", postgressDb)
	prod := queue.NewTransactionProducer(rabbitCh, queueName)
	transactionSvc := txSvc.NewTransactionService(prod, transactionRepo)
	transactionWeb := txWeb.NewTransactionHandler(transactionSvc)

	queue.StartConsumer(rabbitCh, queueName, transactionSvc)

	r.POST("/accounts", accountWeb.CreateAccount)
	r.GET("/accounts/:id", accountWeb.GetAccount)
	r.POST("/transactions", transactionWeb.CreateTransaction)
	r.GET("/accounts/:id/transactions", transactionWeb.GetTransactions)

	return r
}
