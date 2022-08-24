databaseString := postgresql://postgres:postgres@localhost:5432?sslmode=disable

run:
	go run .

build:
	go build .

build_and_run:
	go build .
	./simple-expense

migrate-up:
	migrate -path resources/db/migration -database ${databaseString} -verbose up

migrate-down:
	migrate -path resources/db/migration -database ${databaseString} -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: run build build_and_run migrate-up migrate-down sqlc test
