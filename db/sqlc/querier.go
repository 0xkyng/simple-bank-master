// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
