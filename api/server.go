package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"fmt"
	"net/http"

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

func (s *Server) ListenAndServe() error {
	port := ":3000"
	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe(port, s.myRouter)

	if err != nil {
		return err
	}

	return nil
}
