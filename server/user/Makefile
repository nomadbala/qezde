include .env

url=$(POSTGRES_URL)

migrate-up:
	migrate -database "$(url)" -path ./db/postgres/migrations/ goto $(version)

migrate-down:
	migrate -database "$(url)" -path ./db/postgres/migrations/ goto $(version)

migrate-create:
	migrate create -ext sql -dir ./db/postgres/migrations -seq $(name)

.PHONY: migrate-up migrate-down migrate-create