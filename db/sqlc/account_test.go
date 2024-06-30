package db

import (
	"context"
	"testing"
	"time"

	"github.com/PyMarcus/go_sqlc/util"
	"github.com/stretchr/testify/require"
)

func createAcc(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createAcc(t)
}

func TestGetAccount(t *testing.T) {
	acc := createAcc(t)
	account, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)
	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc := createAcc(t)
	params := UpdateAccountParams{
		ID:      acc.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, acc.ID, account.ID)
	require.NotEqual(t, acc.Balance, account.Balance)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, acc.Currency, account.Currency)
	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T){
	acc := createAcc(t)
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	
	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.Empty(t, acc2)
}
