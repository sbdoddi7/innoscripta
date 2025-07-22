package repository

import (
	"database/sql"
	"fmt"

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

func (ar *accountRepository) CreateAccount(req model.CreateAccountReq) (int64, error) {
	stmt, err := ar.db.Prepare(queryCreateAccount)
	if err != nil {
		return 0, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	var accountNumber int64
	err = stmt.QueryRow(req.FirstName, req.LastName, req.Balance).Scan(&accountNumber)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %w", err)
	}

	return accountNumber, nil
}

func (ar *accountRepository) GetAccount(id string) (model.Account, error) {
	// DB logic
	stmt, err := ar.db.Prepare(queryGetAccountByNumber)
	if err != nil {
		return model.Account{}, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	var account model.Account

	err = stmt.QueryRow(queryGetAccountByNumber).Scan(&account)
	if err != nil {
		return model.Account{}, fmt.Errorf("insert failed: %w", err)
	}
	return account, nil
}
