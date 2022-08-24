package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomInvoice(t *testing.T) Invoice{
	arg := CreateInvoiceParams{
		OrganisationID: CreateRandomOrganisation(t).ID,
		Uploader:       CreateRandomUser(t).ID,
		ExpenseID:      CreateRandomExpense(t).ID,
	}

	invoice, err := testQueries.CreateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)

	require.Equal(t, arg.OrganisationID, invoice.OrganisationID)
	require.Equal(t, arg.Uploader, invoice.Uploader)

	require.NotZero(t, invoice.ID)
	require.NotZero(t, invoice.CreatedAt)

    return invoice
}

func TestCreateRandomInvoice(t *testing.T) {
    CreateRandomInvoice(t)
}
