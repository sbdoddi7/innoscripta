package service

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/sbdoddi7/innoscripta/src/account/mocks"
	"github.com/sbdoddi7/innoscripta/src/model"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountRepository(ctrl)
	service := NewAccountService(mockRepo)

	tests := []struct {
		name      string
		input     model.CreateAccountReq
		mockSetup func()
		wantID    int64
		wantErr   bool
	}{
		{
			name: "success",
			input: model.CreateAccountReq{
				FirstName: "Soma",
				LastName:  "Doddi",
				Balance:   1000,
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					CreateAccount(gomock.Any()).
					Return(int64(123), nil)
			},
			wantID:  123,
			wantErr: false,
		},
		{
			name: "repo error",
			input: model.CreateAccountReq{
				FirstName: "Soma",
				LastName:  "Doddi",
				Balance:   1000,
			},
			mockSetup: func() {
				mockRepo.EXPECT().
					CreateAccount(gomock.Any()).
					Return(int64(0), errors.New("db error"))
			},
			wantID:  0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			gotID, err := service.CreateAccount(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantID, gotID)
			}
		})
	}
}

func TestAccountService_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockAccountRepository(ctrl)
	service := NewAccountService(mockRepo)

	tests := []struct {
		name        string
		inputID     int64
		mockSetup   func()
		wantAccount model.Account
		wantErr     bool
	}{
		{
			name:    "success",
			inputID: 1,
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAccount(int64(1)).
					Return(model.Account{
						ID:        1,
						FirstName: "Soma",
						LastName:  "Doddi",
						Balance:   1000,
						CreatedAt: time.Time{},
					}, nil)
			},
			wantAccount: model.Account{
				ID:        1,
				FirstName: "Soma",
				LastName:  "Doddi",
				Balance:   1000,
				CreatedAt: time.Time{},
			},
			wantErr: false,
		},
		{
			name:    "repo error",
			inputID: 2,
			mockSetup: func() {
				mockRepo.EXPECT().
					GetAccount(int64(2)).
					Return(model.Account{}, errors.New("not found"))
			},
			wantAccount: model.Account{},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			got, err := service.GetAccount(tt.inputID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantAccount, got)
			}
		})
	}
}
