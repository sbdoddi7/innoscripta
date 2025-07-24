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

// CreateAccount inserts a new account into the database with the provided
// first name, last name, and initial balance. It returns the newly generated
// account number if successful.
func (ar *accountRepository) CreateAccount(req model.CreateAccountReq) (int64, error) {
	stmt, err := ar.db.Prepare(queryCreateAccount)
	if err != nil {
		return 0, fmt.Errorf("prepare statement failed: %w", err)
	}
	defer stmt.Close()

	var accountNumber int64
	err = stmt.QueryRow(req.FirstName, req.LastName, req.Balance).Scan(&accountNumber)
	if err != nil {
		return 0, fmt.Errorf("insert account failed: %w", err)
	}

	return accountNumber, nil
}

// GetAccount fetches account details with provided account number from database.
func (ar *accountRepository) GetAccount(id int64) (model.Account, error) {
	stmt, err := ar.db.Prepare(queryGetAccountByNumber)
	if err != nil {
		return model.Account{}, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	var account model.Account

	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return model.Account{}, fmt.Errorf("query account failed: %w", err)
	}
	return account, nil
}
