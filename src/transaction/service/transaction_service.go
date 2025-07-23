package service

import (
	"github.com/google/uuid"
	"github.com/sbdoddi7/innoscripta/src/model"
	"github.com/sbdoddi7/innoscripta/src/platform/queue"
)

type transactionService struct {
	producer queue.TransactionProducer
	repo     model.TransactionRepository // for consumer to write logs
}

func NewTransactionService(p queue.TransactionProducer, repo model.TransactionRepository) *transactionService {
	return &transactionService{producer: p, repo: repo}
}

func (ts *transactionService) CreateTransaction(accountNumber int64, amount float64, txType string) (string, error) {
	txID := uuid.NewString()
	msg := model.TransactionMessage{
		TransactionID: txID,
		AccountNumber: accountNumber,
		Amount:        amount,
		Type:          txType,
	}
	return txID, ts.producer.Publish(msg)
}

// consumer calls this
func (ts *transactionService) ProcessTransaction(msg model.TransactionMessage) error {
	return ts.repo.WriteLog(msg.ToTransactionLog())
}

func (ts *transactionService) GetTransactions(accountNumber int64, page, limit int64) ([]model.TransactionLog, error) {
	offset := (page - 1) * limit
	return ts.repo.GetTransactions(accountNumber, limit, offset)
}
