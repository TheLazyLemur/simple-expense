package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	store *db.Store
}

func NewServer(store *db.Store) *Server {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arg := db.CreateUserParams{
			Name:     util.RandomUsername(),
			Email:    util.RandomEmail(),
			Password: util.RandomString(8),
			Username: util.RandomUsername(),
			Salt:     util.RandomString(10),
		}

		_, _ = store.CreateUser(context.Background(), arg)
	})

	return &Server{store: store}
}

func (s *Server) ListenAndServe() error {
	port := ":3000"
	fmt.Println("Server is listening on port", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		return err
	}

	return nil
}
