//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	connString    string = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
	migrationPath string = "resources/db/migration"
)

// Run with go run .
func Run() error {
	mg.Deps(MigrateUp)

	err := sh.RunV("go", "run", ".")
	if err != nil {
		return err
	}
	return nil
}

// Build with go build .
func Build() error {
	err := sh.RunV("go", "build", ".")
	if err != nil {
		return err
	}
	return nil
}

// Build and run
func BuildAndRun() error {
	err := Build()
	if err != nil {
		return err
	}

	err = sh.RunV("./simple-expense")
	if err != nil {
		return err
	}
	return nil
}

// MigrationUp runs migration up
func MigrateUp() error {
	err := sh.RunV("migrate", "-path", migrationPath, "-database", connString, "-verbose", "up")
	if err != nil {
		return err
	}

	return nil
}

// MigrationDown runs migration down
func MigrateDown() error {
	err := sh.RunV("migrate", "-path", migrationPath, "-database", connString, "-verbose", "down")
	if err != nil {
		return err
	}

	return nil
}

// Generate sqlc go code
func Sqlc() error {
	err := sh.RunV("sqlc", "generate")
	if err != nil {
		return err
	}

	return nil
}

// Clean cache and run tests
func Test() error {
	mg.Deps(MigrateUp)
	err := sh.Run("go", "clean", "-testcache")
	if err != nil {
		return err
	}

	err = sh.Run("go", "test", "-v", "-cover", "./...")
	if err != nil {
		return err
	}
	return nil
}
