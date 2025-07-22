package repository

import (
	"database/sql"

	"github.com/sbdoddi7/innoscripta/src/model"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *accountRepository {
	return &accountRepository{
		db: db,
	}
}

func (ar *accountRepository) CreateAccount() error {
	// DB logic
	return nil
}

func (ar *accountRepository) GetAccount() (model.Account, error) {
	// DB logic
	return model.Account{}, nil
}
