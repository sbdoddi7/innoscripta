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

func (as *accountService) CreateAccount(req model.CreateAccountReq) (int64, error) {
	accountNumber, err := as.repo.CreateAccount(req)
	if err != nil {
		return 0, err
	}
	// DB logic
	return accountNumber, nil
}

func (as *accountService) GetAccount(id string) (model.Account, error) {
	// DB logic
	account, err := as.repo.GetAccount(id)
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
