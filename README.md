## air-verse for hot reload

1. <https://github.com/air-verse/air>
2. run air with `air -c .air.toml`

## Migration

1. Create migration file:
   migrate create -seq -ext sql -dir ././cmd/migrate/migrations create_users

2. Perform migration:
   migrate -path ./cmd/migrate/migrations -database="postgres://postgres:postgres@localhost:5432/go-ecommerce-db?sslmode=disable" up

## Docker

1. running docker container that we specify in docker-compose.yml:
   docker compose up -d

2. stop docker container:
   docker compose down

3. remove docker container with its volumes:
   docker compose down -v

4. if you encounter failed to connect to db, try to stop docker compose and run again, its because the database is not yet created

## .air.toml

1. current working is for linux because we are using docker for running this apps
   bin = "./bin/api"
   cmd = "go build -o ./bin/api ./cmd/api/"

2. if you running locally, change .air.toml line 7-8 to:
   bin = "./bin/api.exe"
   cmd = "go build -o ./bin/ ./cmd/api/"

## db connection on host / local machine

1. use address localhost:5433 for connecting to db

## POSTGRES

1. For now the postgres using extension pg_trgm, if you want to disable the extension, please run this sql command first

Before disabling pg_trgm, remove any indexes using gin_trgm_ops:

DROP INDEX IF EXISTS idx_products_name_trgm;

DROP EXTENSION IF EXISTS pg_trgm CASCADE;
