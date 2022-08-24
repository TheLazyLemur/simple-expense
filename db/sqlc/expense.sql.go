// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: expense.sql

package db

import (
	"context"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (organisation_id, uploader, amount)
VALUES ($1, $2, $3) 
RETURNING id, uploader, amount, organisation_id, created_at
`

type CreateExpenseParams struct {
	OrganisationID int64 `json:"organisation_id"`
	Uploader       int64 `json:"uploader"`
	Amount         int64 `json:"amount"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense, arg.OrganisationID, arg.Uploader, arg.Amount)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Uploader,
		&i.Amount,
		&i.OrganisationID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpense = `-- name: GetExpense :one
SELECT id, uploader, amount, organisation_id, created_at FROM expenses
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExpense(ctx context.Context, id int64) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Uploader,
		&i.Amount,
		&i.OrganisationID,
		&i.CreatedAt,
	)
	return i, err
}

const listExpense = `-- name: ListExpense :many
SELECT id, uploader, amount, organisation_id, created_at FROM expenses
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListExpenseParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListExpense(ctx context.Context, arg ListExpenseParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpense, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Uploader,
			&i.Amount,
			&i.OrganisationID,
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

const listExpenseByOrganisation = `-- name: ListExpenseByOrganisation :many
SELECT id, uploader, amount, organisation_id, created_at FROM expenses
WHERE organisation_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListExpenseByOrganisationParams struct {
	OrganisationID int64 `json:"organisation_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

func (q *Queries) ListExpenseByOrganisation(ctx context.Context, arg ListExpenseByOrganisationParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenseByOrganisation, arg.OrganisationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Uploader,
			&i.Amount,
			&i.OrganisationID,
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
