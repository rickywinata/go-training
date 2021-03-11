# Catalog 4

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
# With an assumption: currently you're on `./catalog4` directory.
#

goose -dir ./internal/postgres/migrations create create_table_product sql
```

## Run migrations

```
# With an assumption: currently you're on `./catalog4` directory
# and you already have database `catalogsvc` exist.
#

goose -dir ./internal/postgres/migrations postgres "user=postgres password=password dbname=catalog sslmode=disable" up
```

## Generate models with sqlboiler.

```
sqlboiler --add-soft-deletes psql
```
