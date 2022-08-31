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

	serv.myRouter.HandleFunc("/login", serv.loginUser).Methods(http.MethodPost)
	serv.myRouter.Handle("/user", ValidateJWT(serv.getUser)).Methods(http.MethodGet)
	serv.myRouter.HandleFunc("/user", serv.newUser).Methods(http.MethodPost)

	return serv
}

func (s *Server) ListenAndServe(port string) error {
	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe(port, s.myRouter)

	if err != nil {
		return err
	}

	return nil
}
