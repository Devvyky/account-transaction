package db

import (
	"context"
	"testing"
	"time"

	"github.com/devvyky/account-transaction/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	documentNumber := util.RandomNumericString(8)

	account, err := testQueries.CreateAccount(context.Background(), documentNumber)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, documentNumber, account.DocumentNumber)
	require.NotZero(t, account.AccountID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.AccountID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.AccountID, account2.AccountID)
	require.Equal(t, account1.DocumentNumber, account2.DocumentNumber)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}
