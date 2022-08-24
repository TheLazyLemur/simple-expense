package db

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Name:  "Dan",
		Email: "danrousseau@protonmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if user.Name != arg.Name {
		t.Errorf("expected user name to be %v. got %v", arg.Name, user.Name)
	}

	if user.Email != arg.Email {
		t.Errorf("expected user email to be %v. got %v", arg.Email, user.Email)
	}
}
