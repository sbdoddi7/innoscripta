package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sbdoddi7/innoscripta/src/model"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_CreateAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	req := model.CreateAccountReq{
		FirstName: "Soma",
		LastName:  "Doddi",
		Balance:   1000,
	}

	// Success case
	mock.ExpectPrepare(regexp.QuoteMeta(queryCreateAccount)).
		ExpectQuery().
		WithArgs(req.FirstName, req.LastName, req.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"account_number"}).AddRow(1))

	id, err := repo.CreateAccount(req)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), id)

	// Error on prepare
	mock.ExpectPrepare("bad query").WillReturnError(assert.AnError)
	badRepo := NewAccountRepository(db)
	_, err = badRepo.CreateAccount(req)
	assert.Error(t, err)
}

func TestAccountRepository_GetAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewAccountRepository(db)

	account := model.Account{
		ID:        1,
		FirstName: "Soma",
		LastName:  "Doddi",
		Balance:   1000,
		CreatedAt: time.Time{}, // or time.Time{} in real match
	}

	// Success case
	mock.ExpectPrepare(regexp.QuoteMeta(queryGetAccountByNumber)).
		ExpectQuery().
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{
			"account_number", "first_name", "last_name", "balance", "created_at",
		}).AddRow(
			account.ID, account.FirstName, account.LastName, account.Balance, account.CreatedAt,
		))

	got, err := repo.GetAccount(1)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, got.ID)
	assert.Equal(t, account.FirstName, got.FirstName)
	assert.Equal(t, account.LastName, got.LastName)
	assert.Equal(t, account.Balance, got.Balance)

	// Error on prepare
	mock.ExpectPrepare("bad query").WillReturnError(assert.AnError)
	badRepo := NewAccountRepository(db)
	_, err = badRepo.GetAccount(1)
	assert.Error(t, err)
}
