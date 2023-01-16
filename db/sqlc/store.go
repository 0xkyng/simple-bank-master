package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to executedb queries and transactions
type Store struct{
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	// Build a newStore object
	// Abd return it
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// Call store.db.BeginTx
	// To begin transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err:%v, rb err:%v", err, rbErr)
		}
	}
	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct{
	FromAccountID	int64 `json:"from_account_id"`
	ToAcountID		int64 `json:"to_account_id"`
	Amount			int64 `json:"amount"`
}

// TransferTxResult the result of the transfer transaction
type TransferTxResult struct{
	Transfer Transfer	`json:"transfer"`
	FromAccount Account	`json:"from_account"`
	ToAccount	Account	`json:"to_account"`
	FromEntry	Entry	`json:"from_entry"`
	ToEntry		Entry	`json:"to_entry"`
}

// TransferTx perfoms a money from one account to the other
// It creates a transfer record, add account entries, 
// And update accounts' balance within a singe db transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	// Create an empty result
	var result TransferTxResult

	// Call store.execTx
	// To create & run a new db transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var  err error

		// Use querie object to call any
		// Individual function that it provides
		result.Transfer, err = q.Createtransfer(ctx, CreateTransferParams{
			FromAccounID:	arg.FromAccountID,
			ToAccountID:	arg.ToAcountID,
			Amount:			arg.Amount,
		})
		if err != nil {
			return err
		}

		// Add FromEntry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount :	-arg.Amount,
		})
		if err != nil {
			return err
		}

		// Add ToEntry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAcountID,
			Amount :	arg.Amount,
		})
		if err != nil {
			return err
		}

		// 


		return nil
	})

	return result, err


}