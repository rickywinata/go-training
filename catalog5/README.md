# Catalog 5

Changelog:
- Add `./internal/catalog/http/error.go`.

## Installing

- Install goose.

```
go get -u github.com/pressly/goose/cmd/goose
```

- Install sqlboiler v4.

```
go get -u -t github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql
```

## Create migrations

```
# With an assumption: currently you're on `./catalogsvc5` directory.
#

goose -dir ./postgres/migrations create create_table_product sql
```

## Run migrations

```
# With an assumption: currently you're on `./catalogsvc5` directory
# and you already have database `catalogsvc` exist.
#

goose -dir ./postgres/migrations postgres "user=postgres password=password dbname=catalogsvc sslmode=disable" up
```

## Generate models with sqlboiler.

```
sqlboiler --add-soft-deletes psql
```
