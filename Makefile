#!make
include cmd/ordersystem/.env
createmigration:
	migrate create -ext=sql -dir sql/migrations -seq init

migrate:
	migrate --path=sql/migrations -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose up

migratedown:
	migrate --path=sql/migrations -database "${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -verbose down

.PHONY: migrate migratedown createmigration