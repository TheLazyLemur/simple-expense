package api

import (
	"TheLazyLemur/simple-expense/auth"
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
	serv.myRouter.Handle("/user", auth.ValidateJWT(serv.getUser)).Methods(http.MethodGet)
	serv.myRouter.HandleFunc("/user", serv.newUser).Methods(http.MethodPost)

	serv.myRouter.Handle("/expense", auth.ValidateJWT(serv.getExpense)).Methods(http.MethodGet)
	serv.myRouter.Handle("/expense", auth.ValidateJWT(serv.newExpense)).Methods(http.MethodPost)

	serv.myRouter.Handle("/invoice", auth.ValidateJWT(serv.newInvoice)).Methods(http.MethodPost)
	serv.myRouter.Handle("/invoice", auth.ValidateJWT(serv.getInvoice)).Methods(http.MethodGet)

	serv.myRouter.Handle("/expense/invoice", auth.ValidateJWT(serv.newExpenseWithInvoice)).Methods(http.MethodPost)

	return serv
}

func (s *Server) ListenAndServe(port string) error {
	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe("0.0.0.0"+":"+port, s.myRouter)

	if err != nil {
		return err
	}

	return nil
}
