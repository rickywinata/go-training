# Team Service

## Installing project dependencies

1. [Go 1.13.0 or newer](https://golang.org/dl/).

1. `goose`.

    ```
    go get -u github.com/pressly/goose/cmd/goose
    ```

1. `sqlboiler` v4.

    ```
    go get -u -t github.com/volatiletech/sqlboiler/v4
    go get -u -t github.com/volatiletech/null/v8
    go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
    ```

## Running locally

These instructions assume that postgres is run & setup with `user=postgres` & `password=password`, and database `teamsvc` exists.

1. Run the migrations

    ```    
    goose -dir ./postgres/migrations postgres "user=postgres password=password dbname=teamsvc sslmode=disable" up
    ```

1. Run `main.go`

    ```
    go run main.go
    ```

## Adding a new database table

1. Create a migration file.

    ```
    goose -dir ./postgres/migration create create_table_team sql
    ```

1. Modify the migration file.

1. Run the migrations.
    ```    
    goose -dir ./postgres/migration postgres "user=postgres password=password dbname=teamsvc sslmode=disable" up
    ```

1. Generate models with sqlboiler.

    ```
    sqlboiler --add-soft-deletes psql
    ```
