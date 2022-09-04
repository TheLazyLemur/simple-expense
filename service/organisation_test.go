package service

import (
	"TheLazyLemur/simple-expense/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganisation(t *testing.T) {
	name, email, username, password := getRandomUserDetails()
    randomOrgName := util.RandomString(10)


	user, err := CreateNewUser(name, email, username, password, store)
    require.NoError(t, err)

    organisation, err := CreateOrganisation(user.ID, randomOrgName, store)

    require.NoError(t, err)
    require.NotEmpty(t, organisation)
    require.Equal(t, organisation.Name, randomOrgName)
    require.Equal(t, organisation.Owner, user.ID)
}

func TestGetOrganisation(t *testing.T) {
	name, email, username, password := getRandomUserDetails()
    randomOrgName := util.RandomString(10)


	user, err := CreateNewUser(name, email, username, password, store)
    require.NoError(t, err)

    organisation, err := CreateOrganisation(user.ID, randomOrgName, store)

    require.NoError(t, err)
    require.NotEmpty(t, organisation)
    require.Equal(t, organisation.Name, randomOrgName)
    require.Equal(t, organisation.Owner, user.ID)

    org, err := GetOrganisation(organisation.ID, store)
    require.NoError(t, err)
    require.NotEmpty(t, org)
    require.Equal(t, org.Name, randomOrgName)
    require.Equal(t, org.Owner, user.ID)
}

func TestAddUserToOrganisation(t *testing.T) {
	name, email, username, password := getRandomUserDetails()
    randomOrgName := util.RandomString(10)


	user, err := CreateNewUser(name, email, username, password, store)
    require.NoError(t, err)

    organisation, err := CreateOrganisation(user.ID, randomOrgName, store)

    require.NoError(t, err)
    require.NotEmpty(t, organisation)
    require.Equal(t, organisation.Name, randomOrgName)
    require.Equal(t, organisation.Owner, user.ID)

	name2, email2, username2, password2 := getRandomUserDetails()
    user2, err := CreateNewUser(name2, email2, username2, password2, store)
    require.NoError(t, err)

    _, err = AddUserToOrganisation(user2.ID, organisation.ID, store)
    require.NoError(t, err)
}

func TestFailAddUserToOrganisation(t *testing.T) {
	name, email, username, password := getRandomUserDetails()
    randomOrgName := util.RandomString(10)


	user, err := CreateNewUser(name, email, username, password, store)
    require.NoError(t, err)

    organisation, err := CreateOrganisation(user.ID, randomOrgName, store)

    require.NoError(t, err)
    require.NotEmpty(t, organisation)
    require.Equal(t, organisation.Name, randomOrgName)
    require.Equal(t, organisation.Owner, user.ID)

	name2, email2, username2, password2 := getRandomUserDetails()
    user2, err := CreateNewUser(name2, email2, username2, password2, store)
    require.NoError(t, err)

    _, err = AddUserToOrganisation(user2.ID, organisation.ID, store)
    require.NoError(t, err)

    _, err = AddUserToOrganisation(user2.ID, organisation.ID, store)
    require.Error(t, err)
    require.Equal(t, err.Error(), "User already exists in organisation")
}

func TestReturnFalseCheckIfUserExistsInOrganisation(t *testing.T) {
	exists, _ := CheckIfUserExistsInOrganisation(500, 500, store)

	require.False(t, exists)
}

func TestReturnTrueCheckIfUserExistsInOrganisation(t *testing.T) {
	name, email, username, password := getRandomUserDetails()
    randomOrgName := util.RandomString(10)


	user, err := CreateNewUser(name, email, username, password, store)
    require.NoError(t, err)

    organisation, err := CreateOrganisation(user.ID, randomOrgName, store)

    require.NoError(t, err)
    require.NotEmpty(t, organisation)
    require.Equal(t, organisation.Name, randomOrgName)
    require.Equal(t, organisation.Owner, user.ID)

   _,  err = AddUserToOrganisation(user.ID, organisation.ID, store)
    require.NoError(t, err)

	exists, _ := CheckIfUserExistsInOrganisation(user.ID, organisation.ID, store)

	require.True(t, exists)
}
