package service

import "github.com/sbdoddi7/innoscripta/src/model"

type accountService struct {
	repo model.AccountRepository
}

func NewAccountService(repo model.AccountRepository) *accountService {
	return &accountService{
		repo: repo,
	}
}

func (as *accountService) CreateAccount() error {
	// DB logic
	return nil
}

func (as *accountService) GetAccount() (model.Account, error) {
	// DB logic
	return model.Account{}, nil
}
