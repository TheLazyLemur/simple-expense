package db

import (
	"context"
	"testing"

	"TheLazyLemur/simple-expense/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomExpense(t *testing.T) Expense {
	arg := CreateExpenseParams{
		Uploader:       CreateRandomUser(t).ID,
		Amount:         util.RandomInt(10, 10000),
		OrganisationID: CreateRandomOrganisation(t).ID,
	}

	expense, err := testQueries.CreateExpense(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.Equal(t, arg.Uploader, expense.Uploader)
	require.Equal(t, arg.Amount, expense.Amount)
	require.Equal(t, arg.OrganisationID, expense.OrganisationID)

	require.NotZero(t, expense.ID)
	require.NotZero(t, expense.CreatedAt)

	return expense
}

func TestCreateExpense(t *testing.T) {
	CreateRandomExpense(t)
}
