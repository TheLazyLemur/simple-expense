package db

import (
	"TheLazyLemur/simple-expense/util"
	"context"
)

// ExpenseWithInvoiceTxParams The parameters required for creating and expense with an invoice upload
type ExpenseWithInvoiceTxParams struct {
	Amount         int64 `json:"amount"`
	UploaderID     int64 `json:"uploader_id"`
	OrganisationID int64 `json:"organisation_id"`
}

// ExpenseWithInvoiceTxResult The result of creating an expense with an invoice upload
type ExpenseWithInvoiceTxResult struct {
	Expense Expense `json:"expense"`
	Invoice Invoice `json:"invoice"`
	Url     string  `json:"url"`
}

// ExpenseWithInvoiceTx The method for creating an expense with an invoice upload
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
