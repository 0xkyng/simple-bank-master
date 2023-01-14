package db

import (
	"context"
	"testing"

	"github.com/codekyng/simple-bank-master/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	// Pass an argument
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// Check that account details 
	// Matches with the input arguments
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency,account.Currency)

	// Check the account id is
	// Automotically genrated by
	// Postgres
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}