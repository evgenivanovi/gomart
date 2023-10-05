## Database

- DSN: `postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable`

### Running PostgresSQL in Docker

```shell
docker run \
  --name gomart-postgres \
  -e POSTGRES_DB=postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  -d postgres:15.3-alpine
```

```shell
docker start shorturl-postgres
```

### Running migration with from GO entrypoint

```shell
export MIGRATION_DIR="./migrations"
export DATABASE_URI="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

go run cmd/migrator/main.go -dir $MIGRATION_DIR -dsn $DATABASE_URI -command up
```

### Running rollback migration with from GO entrypoint

```shell
export MIGRATION_DIR="./migrations"
export DATABASE_URI="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

go run cmd/migrator/main.go -dir $MIGRATION_DIR -dsn $DATABASE_URI -command down
```

### Running JET

```shell
export DATABASE_URI="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

# Will generate following folder structure:
# - ./postgres/public/model
# - ./postgres/public/table
jet -dsn=$DATABASE_URI -path=./internal/repository
```

---
