package db

import (
	"context"
	"testing"

	"github.com/devvyky/account-transaction/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransaction(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateTransactionParams{
		AccountID:       account.AccountID,
		OperationTypeID: 4,
		Amount:          util.RandomAmount(),
	}

	txn1, err := testQueries.CreateTransaction(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, txn1)

	require.Equal(t, arg.AccountID, txn1.AccountID)
	require.Equal(t, arg.OperationTypeID, txn1.OperationTypeID)
	require.Equal(t, arg.Amount, txn1.Amount)

	// for invalid accountID
	arg.AccountID = 0
	txn2, err := testQueries.CreateTransaction(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, txn2)

	// for invalid operation type id
	arg.AccountID = account.AccountID
	arg.OperationTypeID = 0
	txn3, err := testQueries.CreateTransaction(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, txn3)

}
