package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	store    *db.Store
	myRouter *mux.Router
}

func NewServer(s *db.Store) *Server {
	serv := &Server{
		store:    s,
		myRouter: mux.NewRouter().StrictSlash(true),
	}

	serv.myRouter.HandleFunc("/users/{id}", serv.getUser).Methods(http.MethodGet)
	serv.myRouter.HandleFunc("/users", serv.newUser).Methods(http.MethodPost)

	return serv
}

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

	arg := db.CreateUserParams{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
		Username: userReq.Username,
		Salt:     util.RandomString(10),
	}

	user, _ := s.store.CreateUser(context.Background(), arg)
	w.WriteHeader(http.StatusCreated)

    userResp := createUserResponse{
        ID:       user.ID,
        Name: user.Name,
        Email: user.Email,
        Username: user.Username,
    }

	pl, _ := json.Marshal(userResp)
	_, err := w.Write(pl)
	if err != nil {
		return
	}
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}

	user, err := s.store.GetUser(context.Background(), int64(userID))
	if err != nil {
		fmt.Println(err)
	}

    getUserResp := getUserResponse{
        Name: user.Name,
        Email: user.Email,
        Username: user.Username,
    }

	pl, err := json.Marshal(getUserResp)
	if err != nil {
		fmt.Println(err)
	}
	_, err = w.Write(pl)
	if err != nil {
		return
	}
}

func (s *Server) ListenAndServe() error {
	port := ":3000"
	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe(port, s.myRouter)

	if err != nil {
		return err
	}

	return nil
}
