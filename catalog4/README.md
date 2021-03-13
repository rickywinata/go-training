# Catalog 4

## Installing project dependencies

- Install goose.

```
go get -u github.com/pressly/goose/cmd/goose
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
   goose -dir ./internal/catalog/migrations postgres "user=postgres password=password dbname=catalog sslmode=disable" up
   ```

1. Run `main.go`

   ```
   go run ./cmd/catalog
   ```

## Adding a new database table

1. Create a migration file.

   ```
   goose -dir ./internal/catalog/migrations create create_table_name_of_table sql
   ```

1. Write the `CREATE TABLE` query in the migration file.

1. Run the migrations.

   ```
   goose -dir ./internal/catalog/migrations postgres "user=postgres password=password dbname=catalog sslmode=disable" up
   ```

1. Generate models with sqlboiler.

   ```
   (cd ./internal/catalog && sqlboiler --add-soft-deletes psql)
   ```
