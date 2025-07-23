package model

import "time"

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type TransactionService interface {
	CreateTransaction(accountNumber int64, amount float64, txType string) (string, error)
	ProcessTransaction(msg TransactionMessage) error
	GetTransactions(accountNumber int64, page, limit int64) ([]TransactionLog, error)
}

type TransactionRepository interface {
	WriteLog(log TransactionLog) error
	GetTransactions(accountNumber int64, limit, offset int64) ([]TransactionLog, error)
	UpdateBalance(accountNumber int64, delta float64) error
}

type CreateTransactionRequest struct {
	AccountNumber int64   `json:"account_number"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
}

type TransactionLog struct {
	ID            string          `bson:"_id,omitempty"`
	TransactionID string          `bson:"transaction_id"`
	AccountNumber int64           `bson:"account_number"`
	Amount        float64         `bson:"amount"`
	Type          TransactionType `bson:"type"` // "deposit", "withdrawal"
	Timestamp     time.Time       `bson:"timestamp"`
}

type TransactionMessage struct {
	TransactionID string          `json:"transaction_id"`
	AccountNumber int64           `json:"account_number"`
	Amount        float64         `json:"amount"`
	Type          TransactionType `json:"type"`
}

func (m TransactionMessage) ToTransactionLog() TransactionLog {
	return TransactionLog{
		TransactionID: m.TransactionID,
		AccountNumber: m.AccountNumber,
		Amount:        m.Amount,
		Type:          m.Type,
	}
}
