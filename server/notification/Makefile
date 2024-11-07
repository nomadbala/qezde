run:
	go run main.go
build:
	docker build -t notification-service .
up:
	docker-compose up -d
down:
	docker-compose down
restart: down up

.PHONY: run build up down restart