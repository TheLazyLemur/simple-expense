package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"TheLazyLemur/simple-expense/util"

	"crypto/sha256"

	"github.com/stretchr/testify/require"
)

var pepper = util.RandomString(10)

func CreateRandomUser(t *testing.T) User {

	salt := util.RandomString(32)
	hash := sha256.Sum256([]byte(util.RandomString(10) + salt + pepper))

	arg := CreateUserParams{
		Name:     util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: fmt.Sprintf("%x", hash),
		Username: util.RandomUsername(),
		Salt:     salt,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestCreateUserFail(t *testing.T) {
	salt := util.RandomString(32)
	hash := sha256.Sum256([]byte(util.RandomString(10) + salt + pepper))

	username := util.RandomUsername()

	arg := CreateUserParams{
		Name:     util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: fmt.Sprintf("%x", hash),
		Username: username,
		Salt:     salt,
	}

	_, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	_, err = testQueries.CreateUser(context.Background(), arg)
	require.Error(t, err)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestUserLogin(t *testing.T) {
	user1 := CreateRandomUser(t)

	arg := LoginUserParams{
		Username: user1.Username,
		Password: user1.Password,
	}

	user2, err := testQueries.LoginUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	arg := UpdateUserParams{
		ID:   user1.ID,
		Name: util.RandomUsername(),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user2.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
