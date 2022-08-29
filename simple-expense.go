package main

import (
	"TheLazyLemur/simple-expense/api"
	"TheLazyLemur/simple-expense/db/sqlc"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
)

func main() {
	m, err := migrate.New("file://./resources/db/migration", dbSource)
	_ = m.Up()

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
