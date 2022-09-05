package api

import (
	"TheLazyLemur/simple-expense/auth"
	"TheLazyLemur/simple-expense/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) newUser(w http.ResponseWriter, r *http.Request) {
	reqBody, readErr := io.ReadAll(r.Body)
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

	user, err := service.CreateNewUser(userReq.Name, userReq.Email, userReq.Username, userReq.Password, s.store)

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
	id := auth.GetClaimsProperty(r, "id").(float64)

	user, err := service.GetSingleUser(int64(id), s.store)
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

	token, err := service.LoginWithAUsername(logInUserReq.Username, logInUserReq.Password, s.store)
	if err != nil || token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	loginResp := loginUserResponse{
		Token: token,
	}

	pl, err := json.Marshal(loginResp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = w.Write([]byte(pl))
	if err != nil {
		return
	}
}
