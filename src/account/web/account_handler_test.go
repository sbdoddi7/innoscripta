package web

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sbdoddi7/innoscripta/src/account/mocks"
	"github.com/sbdoddi7/innoscripta/src/model"
)

func TestHandler_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mocks.NewMockAccountService(ctrl)

	handler := NewAccountHandler(mockSvc)

	router := gin.Default()
	router.POST("/accounts", handler.CreateAccount)

	tests := []struct {
		name         string
		body         string
		mockSetup    func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "success",
			body: `{"first_name":"Soma","last_name":"Doddi","balance":1000}`,
			mockSetup: func() {
				mockSvc.EXPECT().
					CreateAccount(model.CreateAccountReq{
						FirstName: "Soma",
						LastName:  "Doddi",
						Balance:   1000,
					}).
					Return(int64(1), nil)
			},
			wantCode:     http.StatusCreated,
			wantResponse: `{"account_number":1, "message":"CreateAccount Success!"}`,
		},
		{
			name: "invalid json",
			body: `{"first_name":"John","balance":"oops"}`, // invalid field
			mockSetup: func() {
				//
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"message":"invalid request"}`,
		},
		{
			name: "service error",
			body: `{"first_name":"Jane","last_name":"Doe","balance":200}`,
			mockSetup: func() {
				mockSvc.EXPECT().
					CreateAccount(model.CreateAccountReq{
						FirstName: "Jane",
						LastName:  "Doe",
						Balance:   200,
					}).
					Return(int64(0), assert.AnError)
			},
			wantCode:     http.StatusInternalServerError,
			wantResponse: `{"message":"internal server error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req, _ := http.NewRequest("POST", "/accounts", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.JSONEq(t, tt.wantResponse, w.Body.String())
		})
	}
}

func TestHandler_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := mocks.NewMockAccountService(ctrl)
	handler := NewAccountHandler(mockSvc)

	router := gin.Default()
	router.GET("/accounts/:id", handler.GetAccount)

	tests := []struct {
		name         string
		url          string
		mockSetup    func()
		wantCode     int
		wantResponse string
	}{
		{
			name: "Get account success",
			url:  "/accounts/1",
			mockSetup: func() {
				mockSvc.EXPECT().
					GetAccount(int64(1)).
					Return(model.Account{
						ID:        1,
						FirstName: "Soma",
						LastName:  "Doddi",
						Balance:   1000,
						CreatedAt: time.Time{},
					}, nil)
			},
			wantCode: http.StatusOK,
			wantResponse: `{
		        "id":1,
		        "first_name":"Soma",
		        "last_name":"Doddi",
		        "balance":1000,
		        "created_at":"0001-01-01T00:00:00Z"
		    }`,
		},
		{
			name: "invalid account id",
			url:  "/accounts/abc",
			mockSetup: func() {
				// service should not be called
			},
			wantCode:     http.StatusBadRequest,
			wantResponse: `{"message":"invalid account number"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req, _ := http.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)
			assert.JSONEq(t, tt.wantResponse, w.Body.String())
		})
	}
}
