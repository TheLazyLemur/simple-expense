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

func TestGetExpense(t *testing.T) {
    expense1 := CreateRandomExpense(t)
    expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, expense2)

    require.Equal(t, expense1.ID, expense2.ID)
    require.Equal(t, expense1.Uploader, expense2.Uploader)
    require.Equal(t, expense1.Amount, expense2.Amount)
    require.Equal(t, expense1.OrganisationID, expense2.OrganisationID)
}

func TestListExpense(t *testing.T) {
    for i := 0; i < 10; i++ {
        CreateRandomExpense(t)
    }

    arg := ListExpenseParams{
        Limit:  5,
        Offset: 5,
    }

    expenses, err := testQueries.ListExpense(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, expenses, 5)

    for _, expense := range expenses {
        require.NotEmpty(t, expense)
    }
}

func TestDeleteExpense(t *testing.T) {
    expense1 := CreateRandomExpense(t)
    err := testQueries.DeleteExpense(context.Background(), expense1.ID)
    require.NoError(t, err)

    expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)
    require.Error(t, err)
    require.Empty(t, expense2)
}
