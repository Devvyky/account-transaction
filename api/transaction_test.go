package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/devvyky/account-transaction/db/mock"
	db "github.com/devvyky/account-transaction/db/sqlc"
	"github.com/devvyky/account-transaction/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateTransactionAPI(t *testing.T) {
	account := randomAccount()
	txn := randomTransaction(account.AccountID)
	operationType := randomOperationType()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"account_id":        account.AccountID,
				"operation_type_id": operationType.OperationTypeID,
				"amount":            txn.Amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.AccountID)).
					Times(1).
					Return(account, nil)
				store.EXPECT().
					GetOperationType(gomock.Any(), gomock.Eq(operationType.OperationTypeID)).
					Times(1).
					Return(operationType, nil)
				store.EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					Return(txn, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchTransaction(t, recorder.Body, txn)
			},
		},
		{
			name: "AccountNotFound",
			body: gin.H{
				"account_id":        account.AccountID,
				"operation_type_id": operationType.OperationTypeID,
				"amount":            txn.Amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.AccountID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
				store.EXPECT().
					GetOperationType(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "OperationTypeNotFound",
			body: gin.H{
				"account_id":        account.AccountID,
				"operation_type_id": operationType.OperationTypeID,
				"amount":            txn.Amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.AccountID)).
					Times(1).
					Return(account, nil)
				store.EXPECT().
					GetOperationType(gomock.Any(), gomock.Eq(operationType.OperationTypeID)).
					Times(1).
					Return(db.OperationType{}, sql.ErrNoRows)
				store.EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"account_id":        account.AccountID,
				"operation_type_id": operationType.OperationTypeID,
				"amount":            100,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.AccountID)).
					Times(1).
					Return(account, nil)
				store.EXPECT().
					GetOperationType(gomock.Any(), gomock.Eq(operationType.OperationTypeID)).
					Times(1).
					Return(operationType, nil)
				store.EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Transaction{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidAmount",
			body: gin.H{
				"account_id":        account.AccountID,
				"operation_type_id": operationType.OperationTypeID,
				"amount":            -txn.Amount,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetOperationType(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateTransaction(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomTransaction(accountID int64) db.Transaction {
	return db.Transaction{
		TransactionID:   util.RandomInt(1, 1000),
		OperationTypeID: util.RandomInt(1, 4),
		Amount:          util.RandomFloat64(1.00, 1000.00),
		AccountID:       accountID,
	}
}

func randomOperationType() db.OperationType {
	return db.OperationType{
		OperationTypeID: util.RandomInt(1, 4),
		Description:     util.RandomString(8),
	}
}

func requireBodyMatchTransaction(t *testing.T, body *bytes.Buffer, transaction db.Transaction) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTransaction db.Transaction
	err = json.Unmarshal(data, &gotTransaction)
	require.NoError(t, err)
	require.Equal(t, transaction, gotTransaction)
}
