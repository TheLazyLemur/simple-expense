package main

import (
	"TheLazyLemur/simple-expense/db/sqlc"
	"TheLazyLemur/simple-expense/api"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
)

func main() {
	con, err := sql.Open(dbDriver, dbSource)
    if err != nil {
        log.Fatal("cannot connect to db: ", err)
    }

	store := db.NewStore(con)
    server := api.NewServer(store)
    err = server.ListenAndServe()
    if err != nil {
        log.Fatal("cannot start server: ", err)
    }
}
