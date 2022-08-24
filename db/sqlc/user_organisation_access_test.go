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
        UserID:        user.ID,
        OrganisationID: organisation.ID,
    }

    userOrganisationAccess, err := testQueries.CreateUserOrganisationAccess(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, userOrganisationAccess)
}

func TestDeleteOrganisationAccess(t *testing.T){
    user := CreateRandomUser(t)
    organisation := CreateRandomOrganisation(t)

    arg := CreateUserOrganisationAccessParams{
        UserID:        user.ID,
        OrganisationID: organisation.ID,
    }

    userOrganisationAccess, err := testQueries.CreateUserOrganisationAccess(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, userOrganisationAccess)

    arg2 := DeleteUserOrganiationAccessParams{
        UserID:        user.ID,
        OrganisationID: organisation.ID,
    }

    err = testQueries.DeleteUserOrganiationAccess(context.Background(), arg2)
    require.NoError(t, err)
}
