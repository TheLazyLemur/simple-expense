run:
	go run .

build:
	go build .

build_and_run:
	go build .
	./simple-expense

migrate-up:
	migrate -path resources/db/migration -database "postgresql://postgres:postgres@localhost:5432?sslmode=disable" -verbose up

migrate-down:
	migrate -path resources/db/migration -database "postgresql://postgres:postgres@localhost:5432?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: run build build_and_run migrate-up migrate-down sqlc
