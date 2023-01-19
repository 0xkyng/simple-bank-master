package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	// Call new store
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxParams)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,
			})

			errs <- err // Send err to errs
			results <- result // Send result to results
			// When sending  data to a channel
			// The channel should be on the left 
			// And the data on the right
		}()
	}

	// Check results
	for i := 0; i < n; i++ {
		err := <- errs // errs receive err
		require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result)

		
	}
}