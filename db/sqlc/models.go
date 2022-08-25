// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"
)

type Expense struct {
	ID             int64     `json:"id"`
	Uploader       int64     `json:"uploader"`
	Amount         int64     `json:"amount"`
	OrganisationID int64     `json:"organisation_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Invoice struct {
	ID             int64     `json:"id"`
	OrganisationID int64     `json:"organisation_id"`
	Uploader       int64     `json:"uploader"`
	ExpenseID      int64     `json:"expense_id"`
	Url            string    `json:"url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Organisation struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Owner     int64     `json:"owner"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
}

type UserOrganisationsAccess struct {
	UserID         int64 `json:"user_id"`
	OrganisationID int64 `json:"organisation_id"`
}
