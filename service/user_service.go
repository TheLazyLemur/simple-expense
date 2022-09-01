package service

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"crypto/sha256"
	"fmt"
)

func CreateNewUser(name string, email string, username string, password string, store *db.Store) (db.User, error) {
	salt := util.RandomString(10)
	hash := sha256.Sum256([]byte(password + salt))

	arg := db.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: fmt.Sprintf("%x", hash),
		Username: username,
		Salt:     salt,
	}

	user, err := store.CreateUser(context.Background(), arg)

	return user, err
}

func GetSingleUser(userID int64, store *db.Store) (db.User, error) {
	user, err := store.GetUser(context.Background(), userID)
	return user, err
}

func LoginWithAUsername(username string, password string, store *db.Store) (string, error) {
	user, err := store.GetUserByUsername(context.Background(), username)
	if err != nil {
		return "", err
	}

	salt := user.Salt
	hash := sha256.Sum256([]byte(password + salt))
	hashedPassword := fmt.Sprintf("%x", hash)
	if hashedPassword != user.Password {
		if err != nil {
			return "", err
		}

		return "", fmt.Errorf("Wrong password")
	}

	token, err := GetJWT(user.Email, user.Username, user.ID)
	if err != nil || token == "" {
		return "", fmt.Errorf("Failed to generate token")
	}

	return token, nil
}
