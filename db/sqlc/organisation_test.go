package db

import (
	"context"
	"testing"

	"TheLazyLemur/simple-expense/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomOrganisation(t *testing.T) Organisation {
	user := CreateRandomUser(t)
	arg := CreateOrganisationParams{
		Name:  util.RandomUsername(),
		Owner: user.ID,
	}

	organisation, err := testQueries.CreateOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, organisation)

	require.Equal(t, arg.Name, organisation.Name)

	require.NotZero(t, organisation.ID)
	require.NotZero(t, organisation.CreatedAt)
	require.NotZero(t, organisation.UpdatedAt)

	return organisation
}

func TestCreateOrganisation(t *testing.T) {
	CreateRandomOrganisation(t)
}

func TestGetOrganisation(t *testing.T) {
	organisation1 := CreateRandomOrganisation(t)

	organisation2, err := testQueries.GetOrganisation(context.Background(), organisation1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, organisation2)

	require.Equal(t, organisation1.ID, organisation2.ID)
	require.Equal(t, organisation1.Name, organisation2.Name)
	require.Equal(t, organisation1.Owner, organisation2.Owner)
	require.Equal(t, organisation1.CreatedAt, organisation2.CreatedAt)
	require.Equal(t, organisation1.UpdatedAt, organisation2.UpdatedAt)
}

func TestUpdateOrganisation(t *testing.T) {
	organisation1 := CreateRandomOrganisation(t)

	arg := UpdateOrganisationParams{
		ID:   organisation1.ID,
		Name: util.RandomUsername(),
	}

	organisation2, err := testQueries.UpdateOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, organisation2)

	require.Equal(t, organisation1.ID, organisation2.ID)
	require.Equal(t, arg.Name, organisation2.Name)
	require.Equal(t, organisation1.Owner, organisation2.Owner)
}

func TestDeleteOrganisation(t *testing.T) {
	organisation1 := CreateRandomOrganisation(t)

	err := testQueries.DeleteOrganisation(context.Background(), organisation1.ID)
	require.NoError(t, err)

	organisation2, err := testQueries.GetOrganisation(context.Background(), organisation1.ID)
	require.Error(t, err)
	require.Empty(t, organisation2)
}

func TestListOrganisations(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomOrganisation(t)
	}

	arg := ListOrganisationParams{
		Limit:  5,
		Offset: 5,
	}

	organisations, err := testQueries.ListOrganisation(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, organisations, 5)

	for _, organisation := range organisations {
		require.NotEmpty(t, organisation)
	}
}
