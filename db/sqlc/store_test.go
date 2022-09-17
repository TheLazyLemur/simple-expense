package db

import (
	"TheLazyLemur/simple-expense/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpenseWithInvoiceTx(t *testing.T) {
	store := NewStore(testDB)

	user := CreateRandomUser(t)

	params := ExpenseWithInvoiceTxParams{
		Owner:  user.ID,
		Amount: util.RandomInt(10, 10000),
	}

	errs := make(chan error)
	results := make(chan ExpenseWithInvoiceTxResult)
	n := 5
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.ExpenseWithInvoiceTx(context.Background(), params)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		res := <-results
		require.NoError(t, err)

		require.NotEmpty(t, res.Invoice.Url)

		require.Equal(t, params.Owner, res.Expense.Owner)
		require.Equal(t, params.Amount, res.Expense.Amount)

		require.Equal(t, params.Owner, res.Invoice.Owner)
		require.Equal(t, res.Expense.ID, res.Invoice.ExpenseID)

		require.Equal(t, res.Expense.Owner, res.Invoice.Owner)
		require.Equal(t, res.Expense.ID, res.Invoice.ExpenseID)
	}

}
