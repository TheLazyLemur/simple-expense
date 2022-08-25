# Simple Expense

Simple Expense is a simple expense tracking API can be used to track expenses as well as invoices and generate reports of you spending habbits

## Dependencies

- [Migrate](https://github.com/golang-migrate/migrate) is used for database migrations
- [Sqlc](https://github.com/kyleconroy/sqlc) is used for generating the database access layer

## Getting started

After cloning the repo run `go mod tidy` to install the dependencies.
For running tests you will need to have a local postgres instance running on port 5432.

### Building

Simple Expense uses the [mage](https://magefile.org/) build tool.

You can run `mage` at the root of the project to see a list of available targets.

