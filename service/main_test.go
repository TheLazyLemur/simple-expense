package service

import (
	db "TheLazyLemur/simple-expense/db/sqlc"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
)

var testDB *sql.DB
var store *db.Store

func TestMain(m *testing.M) {
	runMigrations()

	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store = db.NewStore(testDB)

	os.Exit(m.Run())
}

func runMigrations() {
	m, err := migrate.New(
		"file://./../resources/db/migration",
		"postgresql://postgres:postgres@localhost:5432?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	_ = m.Up()
}
