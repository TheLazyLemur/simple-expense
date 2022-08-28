package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	email2 := RandomEmail()

	require.NotEmpty(t, email)
	require.NotEmpty(t, email2)
	require.NotEqual(t, email, email2)
}

func TestRandomString(t *testing.T) {
	str := RandomString(10)
	str2 := RandomString(10)

	require.NotEmpty(t, str)
	require.NotEmpty(t, str2)
	require.NotEqual(t, str, str2)
}

func TestRandomUsername(t *testing.T) {
	username := RandomUsername()
	username2 := RandomUsername()

	require.NotEmpty(t, username)
	require.NotEmpty(t, username2)
	require.NotEqual(t, username, username2)
}

func TestRandomInt(t *testing.T) {
	min := int64(1)
	max := int64(10000)

	num := RandomInt(min, max)
	num2 := RandomInt(min, max)

	require.NotEmpty(t, num)
	require.NotEmpty(t, num2)
	require.NotEqual(t, num, num2)
}
