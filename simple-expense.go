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
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	dbString := os.Getenv("dbSource")
	if dbString != "" {
		dbSource = dbString
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	m, err := migrate.New("file://./resources/db/migration", dbSource)
	if err != nil {
		log.Fatal(err)
	}
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
