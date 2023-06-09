// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const deleteTransfersByFromAccount = `-- name: DeleteTransfersByFromAccount :exec
DELETE FROM transfers WHERE from_account_id = $1
`

func (q *Queries) DeleteTransfersByFromAccount(ctx context.Context, fromAccountID int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfersByFromAccount, fromAccountID)
	return err
}

const deleteTransfersByFromAccountAndToAccount = `-- name: DeleteTransfersByFromAccountAndToAccount :exec
DELETE FROM transfers WHERE from_account_id = $1 AND to_account_id = $2
`

type DeleteTransfersByFromAccountAndToAccountParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
}

func (q *Queries) DeleteTransfersByFromAccountAndToAccount(ctx context.Context, arg DeleteTransfersByFromAccountAndToAccountParams) error {
	_, err := q.db.ExecContext(ctx, deleteTransfersByFromAccountAndToAccount, arg.FromAccountID, arg.ToAccountID)
	return err
}

const deleteTransfersByToAccount = `-- name: DeleteTransfersByToAccount :exec
DELETE FROM transfers WHERE to_account_id = $1
`

func (q *Queries) DeleteTransfersByToAccount(ctx context.Context, toAccountID int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfersByToAccount, toAccountID)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfer = `-- name: ListTransfer :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTransferParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByFromAccount = `-- name: ListTransfersByFromAccount :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 
LIMIT $2
OFFSET $3
`

type ListTransfersByFromAccountParams struct {
	FromAccountID int64 `json:"from_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfersByFromAccount(ctx context.Context, arg ListTransfersByFromAccountParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByFromAccount, arg.FromAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByFromAndToAccount = `-- name: ListTransfersByFromAndToAccount :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2 
LIMIT $3
OFFSET $4
`

type ListTransfersByFromAndToAccountParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfersByFromAndToAccount(ctx context.Context, arg ListTransfersByFromAndToAccountParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByFromAndToAccount,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByToAccount = `-- name: ListTransfersByToAccount :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE to_account_id = $1 
LIMIT $2
OFFSET $3
`

type ListTransfersByToAccountParams struct {
	ToAccountID int64 `json:"to_account_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) ListTransfersByToAccount(ctx context.Context, arg ListTransfersByToAccountParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByToAccount, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfer = `-- name: UpdateTransfer :one
UPDATE transfers SET amount = $2
WHERE id = $1
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type UpdateTransferParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransfer, arg.ID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
