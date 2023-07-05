## Dependencies

Docker (windows only for running sqlc)

[Sqlc](https://sqlc.dev/)
Install sqlc

```sh
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

[Goose](https://github.com/pressly/goose)

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

## Helpful links for goose

[Using Goose CLI](https://citizix.com/managing-database-migrations-with-golang-goose-using-incremental-sql-changes/)
