export

DB_MIGRATE_URL = postgres://login:pass@localhost:5432/postgres?sslmode=disable
MIGRATE_PATH = ./migration/postgres
CONFIG_PATH=.env.local

mod:
	go mod tidy

dev: mod
	CONFIG_PATH="${CONFIG_PATH}" go run ./cmd/app

start:
	go build -o ./bin/app ./cmd/app && ./bin/app

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $(name)

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all

migrate-reset: migrate-down migrate-up