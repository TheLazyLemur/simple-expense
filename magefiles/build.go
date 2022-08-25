//go:build mage

package main

import (
    "github.com/magefile/mage/sh"
)

var (
    connString string = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
    migrationPath string = "resources/db/migration"                                    
)

// Run with go run .
func Run() error {
    err := sh.Run("go", "run", ".")
    if err != nil {
        return err
    }
    return nil
}

// Build with go build .
func Build() error {
    err := sh.Run("go", "build", ".")
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

    err = sh.Run("./simple-expense")
    if err != nil {
        return err
    }
    return nil
}

// MigrationUp runs migration up
func MigrateUp() error {
    err := sh.Run("migrate", "-path", migrationPath, "-database", connString, "-verbose", "up")
    if err != nil {
        return err
    }

    return nil
}

// MigrationDown runs migration down
func MigrateDown() error {
    err := sh.Run("migrate", "-path", migrationPath, "-database", connString, "-verbose", "down")
    if err != nil {
        return err
    }

    return nil
}

// Generate sqlc go code
func Sqlc() error {
    err := sh.Run("sqlc", "generate")
    if err != nil {
        return err
    }

    return nil
}

// Clean cache and run tests
func Test() error {
    err := sh.Run("go", "clean", "-testcache")
    if err != nil {
        return err
    }

    err = sh.Run("go", "test", "-v", "./...")
    if err != nil {
        return err
    }
    return nil
}
