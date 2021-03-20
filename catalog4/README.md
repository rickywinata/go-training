# Catalog 4

## Installing project dependencies

- Install golang-migrate.

```
https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md#installation
```

- Install sqlboiler v4.

```
go get -u -t github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
```

## Running locally

These instructions assume that postgres is run & setup with `user=postgres` & `password=password`, and database `catalog` exists.

1. Run the migrations

   ```
   migrate -source file://internal/catalog/migrations -database "postgres://postgres:password@localhost:5432/catalog?sslmode=disable" up
   ```

1. Run `main.go`

   ```
   go run ./cmd/catalog
   ```

## Adding a new database table

1. Create a migration file.

   ```
   migrate create -ext sql -dir internal/catalog/migrations -seq create_table_name
   ```

1. Write the `CREATE TABLE` query in the `.up.sql` migration file.

1. Run the migrations.

   ```
   migrate -source file://internal/catalog/migrations -database "postgres://postgres:password@localhost:5432/catalog?sslmode=disable" up
   ```

1. Generate models with sqlboiler.

   ```
   (cd ./internal/catalog && sqlboiler --add-soft-deletes psql)
   ```
