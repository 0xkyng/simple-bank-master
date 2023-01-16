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