# Catalog 3

## Requirements

- Go 1.16 or above
- Docker
- `golang-migrate`

   ```
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```


## Running locally

These instructions assume that postgres is run & setup with `user=postgres` & `password=password`, and database `catalog` exists.

1. Run the migrations

   ```
   migrate -source file://internal/database/migrations -database "postgres://postgres:password@localhost:5432/catalog?sslmode=disable" up
   ```

1. Run `main.go`

   ```
   go run ./cmd/catalog
   ```

## Adding a new database table

1. Create a migration file.

   ```
   migrate create -ext sql -dir internal/catalog/database/migrations -seq create_table_name
   ```

1. Write the `CREATE TABLE` query in the `.up.sql` migration file.

1. Run the migrations.

   ```
   migrate -source file://internal/catalog/database/migrations -database "postgres://postgres:password@localhost:5432/catalog?sslmode=disable" up
   ```

1. Generate models with sqlboiler.

   ```
   (cd ./internal/catalog && sqlboiler --add-soft-deletes psql)
   ```
