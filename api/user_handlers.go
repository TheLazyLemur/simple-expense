package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) newUser(w http.ResponseWriter, r *http.Request) {

	reqBody, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return
	}

	var userReq createUserRequest
	jsonErr := json.Unmarshal(reqBody, &userReq)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return
	}

	salt := util.RandomString(10)
	hash := sha256.Sum256([]byte(userReq.Password + salt))

	arg := db.CreateUserParams{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: fmt.Sprintf("%x", hash),
		Username: userReq.Username,
		Salt:     salt,
	}

	user, err := s.store.CreateUser(context.Background(), arg)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)

	userResp := createUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	pl, _ := json.Marshal(userResp)
	_, err = w.Write(pl)
	if err != nil {
		return
	}
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {

    token := r.Header.Get("Token")
    claims, err := DecodeJwt(token)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    id := claims["id"].(float64)

	user, err := s.store.GetUser(context.Background(), int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getUserResp := getUserResponse{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	pl, err := json.Marshal(getUserResp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = w.Write(pl)
	if err != nil {
		return
	}
}

func (s *Server) loginUser(w http.ResponseWriter, r *http.Request) {
	logInUserReq := loginUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&logInUserReq)
	if err != nil {
		log.Fatal(err)
	}

	username := logInUserReq.Username
	password := logInUserReq.Password

	user, err := s.store.GetUserByUsername(context.Background(), username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	salt := user.Salt
	hash := sha256.Sum256([]byte(password + salt))
	hashedPassword := fmt.Sprintf("%x", hash)
	if hashedPassword != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
		return
	}

    token, err := GetJWT(user.Email, user.Username, user.ID)
    if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
    }

    w.Write([]byte(token))
}
