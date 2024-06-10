// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  document_number
) VALUES (
  $1
) RETURNING account_id, document_number, created_at
`

func (q *Queries) CreateAccount(ctx context.Context, documentNumber string) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, documentNumber)
	var i Account
	err := row.Scan(&i.AccountID, &i.DocumentNumber, &i.CreatedAt)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT account_id, document_number, created_at FROM accounts
WHERE account_id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, accountID int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, accountID)
	var i Account
	err := row.Scan(&i.AccountID, &i.DocumentNumber, &i.CreatedAt)
	return i, err
}