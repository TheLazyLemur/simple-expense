package service

import (
	"TheLazyLemur/simple-expense/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func getRandomUserDetails() (string, string, string, string) {
	name := util.RandomUsername()
	email := util.RandomEmail()
	password := util.RandomString(10)
	username := util.RandomUsername()

	return name, email, username, password
}

func TestCreateNewUser(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	user, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.Equal(t, username, user.Username)
	require.NotEmpty(t, user.Salt)
}

func TestCreateNewUserWithExistingUsername(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	_, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	_, err = CreateNewUser(name, email, username, password, store)
	require.Error(t, err)
}

func TestLoginUserWithUsernameAndPassword(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	_, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	token, err := LoginWithAUsername(username, password, store)

	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestFailLoginUserWithWrongPassword(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	_, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	token, err := LoginWithAUsername(username, "wrong password", store)

	require.Error(t, err)
	require.Empty(t, token)
}

func TestFailLoginUserWithWrongUsername(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	_, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	token, err := LoginWithAUsername("Wrong Username", password, store)

	require.Error(t, err)
	require.Empty(t, token)
}

func TestGetSingleUser(t *testing.T) {
	name, email, username, password := getRandomUserDetails()

	user, err := CreateNewUser(name, email, username, password, store)
	require.NoError(t, err)

	user, err = GetSingleUser(user.ID, store)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.Equal(t, username, user.Username)
	require.NotEmpty(t, user.Salt)
}
