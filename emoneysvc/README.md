# E-Money Service

## Installing project dependencies

1. [Go 1.13.0 or newer](https://golang.org/dl/).

1. `goose`.

    ```
    go get -u github.com/pressly/goose/cmd/goose
    ```

1. `sqlboiler` v4.

    ```
    go get -u -t github.com/volatiletech/sqlboiler/v4
    go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
    ```

## Running locally

These instructions assume that postgres is run & setup with `user=postgres` & `password=postgres`, and database `emoneysvc` exists.

1. Run the migrations

    ```    
    goose -dir ./postgres/migrations postgres "user=postgres password=password dbname=emoneysvc sslmode=disable" up
    ```

2. Run `main.go`

    ```
    go run main.go
    ```

## Adding a new database table

1. Create migrations

    ```
    goose -dir ./postgres/migrations create create_table_account sql
    ```


2. Generate models with sqlboiler.

    ```
    sqlboiler --add-soft-deletes psql
    ```
