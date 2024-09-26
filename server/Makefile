SERVICE_NAMES = api-gateway-service-app

build:
	@echo "Building Docker images..."
	@for service in $(SERVICE_NAMES); do \
		docker build -t $$service ./$$service; \
	done
up:
	docker-compose up -d
up_with_build:
	docker-compose up --build
down:
	docker-compose down
restart: down up

.PHONY: build up up_with_build down restart