package db

import (
	"context"
	"database/sql"
	"testing"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Name: sql.NullString{
			String: "Dan",
			Valid:  true,
		},
		Email: sql.NullString{
			String: "danrousseau@protonmail.com",
			Valid:  true,
		},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if user.Name.String != arg.Name.String {
		t.Errorf("expected user name to be %v. got %v", arg.Name, user.Name)
	}

	if user.Email.String != arg.Email.String {
		t.Errorf("expected user email to be %v. got %v", arg.Email, user.Email)
	}
}
