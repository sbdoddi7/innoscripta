package model

import "time"

type Account struct {
	ID        string    `json:"id" db:"id"`
	OwnerName string    `json:"owner_name" db:"owner_name"`
	Balance   float64   `json:"balance" db:"balance"`
	Currency  string    `json:"currency" db:"currency"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type AccountService interface {
	CreateAccount() error
	GetAccount() (Account, error)
}

type AccountRepository interface {
	CreateAccount() error
	GetAccount() (Account, error)
}
