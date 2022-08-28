package db

import (
	"TheLazyLemur/simple-expense/util"
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type ExpenseWithInvoiceTxParams struct {
	Amount         int64 `json:"amount"`
	UploaderID     int64 `json:"uploader_id"`
	OrganisationID int64 `json:"organisation_id"`
}

type ExpenseWithInvoiceTxResult struct {
	Expense Expense `json:"expense"`
	Invoice Invoice `json:"invoice"`
	Url     string  `json:"url"`
}

func (store *Store) ExpenseWithInvoiceTx(ctx context.Context, arg ExpenseWithInvoiceTxParams) (ExpenseWithInvoiceTxResult, error) {
	var result ExpenseWithInvoiceTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Expense, err = q.CreateExpense(ctx, CreateExpenseParams{
			Amount:         arg.Amount,
			Uploader:       arg.UploaderID,
			OrganisationID: arg.OrganisationID,
		})
		if err != nil {
			return err
		}

		result.Invoice, err = q.CreateInvoice(ctx, CreateInvoiceParams{
			ExpenseID:      result.Expense.ID,
			Uploader:       arg.UploaderID,
			OrganisationID: arg.OrganisationID,
			Url:            util.RandomString(10),
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
