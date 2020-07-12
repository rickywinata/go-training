# E-Money Service

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
# With an assumption: currently you're on `./emoneysvc` directory.
#

goose -dir ./postgres/migrations create create_table_account sql
```

## Run migrations

```
# With an assumption: currently you're on `./emoneysvc` directory
# and you already have database `emoneysvc` exist.
#

goose -dir ./postgres/migrations postgres "user=postgres password=password dbname=emoneysvc sslmode=disable" up
```

## Generate models with sqlboiler.

```
sqlboiler --add-soft-deletes psql
```
