package db

import (
	"TheLazyLemur/simple-expense/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomInvoice(t *testing.T) Invoice {
	arg := CreateInvoiceParams{
		Owner:     CreateRandomUser(t).ID,
		ExpenseID: CreateRandomExpense(t).ID,
        Url: util.RandomString(10),
	}

	invoice, err := testQueries.CreateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)

	require.Equal(t, arg.Owner, invoice.Owner)
    require.Equal(t, arg.ExpenseID, invoice.ExpenseID)
    require.Equal(t, arg.Url, invoice.Url)

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
	require.Equal(t, invoice1.Owner, invoice2.Owner)
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
