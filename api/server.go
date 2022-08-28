package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"encoding/json"
	"fmt"
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

	serv.myRouter.HandleFunc("/users/{id}", serv.getUser).Methods("GET")
	serv.myRouter.HandleFunc("/users", serv.newUser).Methods("POST")

	return serv
}

func (s *Server) newUser(w http.ResponseWriter, r *http.Request) {
	arg := db.CreateUserParams{
		Name:     util.RandomUsername(),
		Email:    util.RandomEmail(),
		Password: util.RandomString(8),
		Username: util.RandomUsername(),
		Salt:     util.RandomString(10),
	}

	user, _ := s.store.CreateUser(context.Background(), arg)

	pl, _ := json.Marshal(user)
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

	pl, err := json.Marshal(user)
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
