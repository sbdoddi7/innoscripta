package model

import "time"

type Account struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Balance   float64   `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateAccountReq struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Balance   float64 `json:"balance"`
}

type AccountService interface {
	CreateAccount(req CreateAccountReq) (int64, error)
	GetAccount(id int64) (Account, error)
}

type AccountRepository interface {
	CreateAccount(req CreateAccountReq) (int64, error)
	GetAccount(id int64) (Account, error)
}
