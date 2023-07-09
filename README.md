## Dependencies

Docker (windows only for running sqlc)

[Sqlc](https://sqlc.dev/)

Install sqlc

```sh
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

[Goose](https://github.com/pressly/goose)

Install Goose

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Docker management

Start

```sh
docker-compose up -d
```

Stop

```sh
docker-compose down
```

## Using sqlc from Windows

Run the following from the root of the project
using a `cmd` shell

```sh
docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
```

## Applying Changes with Goose

Create migrations

```sh
goose -dir ./sql/schema create migration_name sql
```

Create seed

```sh
goose -dir ./sql/seeds create seed_name sql
```

Run migrations

```sh
goose -dir ./sql/schema {up|down}
```

Run Seeds

```sh
goose -dir ./schema/seed -no-versioning up
```

[Using Goose CLI](https://citizix.com/managing-database-migrations-with-golang-goose-using-incremental-sql-changes/)
