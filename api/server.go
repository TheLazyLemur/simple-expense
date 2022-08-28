package api

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	store *db.Store
}

func NewServer(store *db.Store) *Server {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			arg := db.CreateUserParams{
				Name:     util.RandomUsername(),
				Email:    util.RandomEmail(),
				Password: util.RandomString(8),
				Username: util.RandomUsername(),
				Salt:     util.RandomString(10),
			}

			user, _ := store.CreateUser(context.Background(), arg)

			// convert user to json
			pl, _ := json.Marshal(user)
			_, err := w.Write(pl)
			if err != nil {
				return
			}
		}

		if r.Method == http.MethodGet {
			user, err := store.GetUser(context.Background(), 127)
			if err != nil {
				fmt.Println(err)
			}

			// convert user to json
			pl, _ := json.Marshal(user)
			_, err = w.Write(pl)
			if err != nil {
				return
			}
		}

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
