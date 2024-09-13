# Golang Simple POS APP

# Description
This is golang simple POS APP.

## Tech Stack

- Golang : https://github.com/golang/go
- SQLite (Database) : https://github.com/sqlite/sqlite

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.yaml` file.

## API Spec

All API Spec is in `api-spec.yaml` file.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```bash
migrate -database sqlite3://pos.db -path db/migrations up
```

```bash
migrate -database sqlite3://pos.db -path db/migrations down
```

## Run Application

### Run unit test

```bash
go test -v ./test/
```

### Run web server

```bash
go run cmd/main.go
```