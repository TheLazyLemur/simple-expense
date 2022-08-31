package main

import (
	"TheLazyLemur/simple-expense/api"
	"TheLazyLemur/simple-expense/db/sqlc"
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

const (
	dbDriver = "postgres"
)

var (
	dbSource = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
)

func main() {
	godotenv.Load()
	dbString := os.Getenv("dbSource")
	if dbString != "" {
		dbSource = dbString
	}

	port := os.Getenv("simpleExpensePort")

	m, err := migrate.New("file://./resources/db/migration", dbSource)
	_ = m.Up()

	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(con)
	server := api.NewServer(store)
	err = server.ListenAndServe(port)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
