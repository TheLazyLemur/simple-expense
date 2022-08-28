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
	org := CreateRandomOrganisation(t)

	params := ExpenseWithInvoiceTxParams{
		UploaderID:     user.ID,
		Amount:         util.RandomInt(10, 10000),
		OrganisationID: org.ID,
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

		require.Equal(t, params.UploaderID, res.Expense.Uploader)
        require.Equal(t, params.Amount, res.Expense.Amount)
        require.Equal(t, params.OrganisationID, res.Expense.OrganisationID)

        require.Equal(t, params.UploaderID, res.Invoice.Uploader)
        require.Equal(t, params.OrganisationID, res.Invoice.OrganisationID)
        require.Equal(t, res.Expense.ID, res.Invoice.ExpenseID)

        require.Equal(t, res.Expense.OrganisationID, res.Invoice.OrganisationID)
        require.Equal(t, res.Expense.Uploader, res.Invoice.Uploader)
        require.Equal(t, res.Expense.ID, res.Invoice.ExpenseID)
	}

}
