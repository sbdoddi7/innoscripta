package service

import (
	"github.com/sbdoddi7/innoscripta/src/model"
)

type accountService struct {
	repo model.AccountRepository
}

func NewAccountService(repo model.AccountRepository) *accountService {
	return &accountService{
		repo: repo,
	}
}

// CreateAccount handles business logic for creating a new account.
// calls the repository to insert the account.
func (as *accountService) CreateAccount(req model.CreateAccountReq) (int64, error) {
	accountNumber, err := as.repo.CreateAccount(req)
	if err != nil {
		return 0, err
	}
	return accountNumber, nil
}

// GetAccount handles business logic for fetching account details
func (as *accountService) GetAccount(id int64) (model.Account, error) {
	account, err := as.repo.GetAccount(id)
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
