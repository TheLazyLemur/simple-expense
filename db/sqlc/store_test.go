package db

import (
	"TheLazyLemur/simple-expense/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpenseWithInvoiceTx(t *testing.T) {
	store := NewStore(testDB)

	// Create a new User
	user := CreateRandomUser(t)

	// Create a new Organisation
	org := CreateRandomOrganisation(t)

	params := ExpenseWithInvoiceTxParams{
		UploaderID:     user.ID,
		Amount:         util.RandomInt(10, 10000),
		OrganisationID: org.ID,
	}

	errs := make(chan error)
	results := make(chan ExpenseWithInvoiceTxResult)

	go func() {
		result, err := store.ExpenseWithInvoiceTx(context.Background(), params)

		errs <- err
		results <- result
	}()

	for i := 0; i < 1; i++ {
		err := <-errs
		require.NoError(t, err)
	}

}
