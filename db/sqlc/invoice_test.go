package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomInvoice(t *testing.T) Invoice {
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

func TestCreateInvoice(t *testing.T) {
	CreateRandomInvoice(t)
}

func TestGetInvoice(t *testing.T) {
    invoice1 := CreateRandomInvoice(t)
    invoice2, err := testQueries.GetInvoice(context.Background(), invoice1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, invoice2)

    require.Equal(t, invoice1.ID, invoice2.ID)
    require.Equal(t, invoice1.OrganisationID, invoice2.OrganisationID)
    require.Equal(t, invoice1.Uploader, invoice2.Uploader)
    require.Equal(t, invoice1.ExpenseID, invoice2.ExpenseID)
}

func TestListInvoice(t *testing.T) {
    for i := 0; i < 10; i++ {
        CreateRandomInvoice(t)
    }

    arg := ListInvoiceParams{
        Limit:  5,
        Offset: 5,
    }

    invoices, err := testQueries.ListInvoice(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, invoices, 5)

    for _, invoice := range invoices {
        require.NotEmpty(t, invoice)
    }
}

func TestDeleteInvoice(t *testing.T) {
    invoice1 := CreateRandomInvoice(t)

    err := testQueries.DeleteInvoice(context.Background(), invoice1.ID)
    require.NoError(t, err)

    invoice2, err := testQueries.GetInvoice(context.Background(), invoice1.ID)
    require.Error(t, err)
    require.Empty(t, invoice2)
}
