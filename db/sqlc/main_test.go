package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	runMigrations()

	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func runMigrations() {
	m, err := migrate.New(
		"file://./../../resources/db/migration",
		"postgresql://postgres:postgres@localhost:5432?sslmode=disable")
	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
