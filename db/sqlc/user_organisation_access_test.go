package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserOrganisationAccess(t *testing.T) {
	user := CreateRandomUser(t)
	organisation := CreateRandomOrganisation(t)

	arg := CreateUserOrganisationAccessParams{
		UserID:         user.ID,
		OrganisationID: organisation.ID,
	}

	userOrganisationAccess, err := testQueries.CreateUserOrganisationAccess(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userOrganisationAccess)
}

func TestGetUserOrganisationAccess(t *testing.T) {
	user := CreateRandomUser(t)
	organisation := CreateRandomOrganisation(t)

	arg := CreateUserOrganisationAccessParams{
		UserID:         user.ID,
		OrganisationID: organisation.ID,
	}

	userOrganisationAccess, err := testQueries.CreateUserOrganisationAccess(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userOrganisationAccess)

	arg2 := GetUserOrganisationAccessParams{
		UserID:         user.ID,
		OrganisationID: organisation.ID,
	}

	userOrganisationAccess2, err := testQueries.GetUserOrganisationAccess(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, userOrganisationAccess2)
}

func TestDeleteOrganisationAccess(t *testing.T) {
	user := CreateRandomUser(t)
	organisation := CreateRandomOrganisation(t)

	arg := CreateUserOrganisationAccessParams{
		UserID:         user.ID,
		OrganisationID: organisation.ID,
	}

	userOrganisationAccess, err := testQueries.CreateUserOrganisationAccess(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, userOrganisationAccess)

	arg2 := DeleteUserOrganisationAccessParams{
		UserID:         user.ID,
		OrganisationID: organisation.ID,
	}

	err = testQueries.DeleteUserOrganisationAccess(context.Background(), arg2)
	require.NoError(t, err)
}
