include .env

IMAGE_NAME = notification-service

DOCKER_DEV_FILE = docker/dev.Dockerfile
DOCKER_PROD_FILE = docker/prod.Dockerfile

DEV_PORT = $(DEV_PORT)

build-dev:
	docker build -f $(DOCKER_DEV_FILE) -t $(IMAGE_NAME)-dev .

build-prod:
	docker build -f $(DOCKER_PROD_FILE) -t $(IMAGE_NAME)-prod .

run-dev: build-dev
	docker run --rm -p $(DEV_PORT):$(DEV_PORT) $(IMAGE_NAME)-dev

run-prod: build-prod
	docker run --rm -p 80:80 $(IMAGE_NAME)-prod

restart-dev:
	docker stop $(shell docker ps -q --filter ancestor=$(IMAGE_NAME)-dev) || true
	docker run --rm -p $(DEV_PORT):$(DEV_PORT) $(IMAGE_NAME)-dev

restart-prod:
	docker stop $(shell docker ps -q --filter ancestor=$(IMAGE_NAME)-prod) || true
	docker run --rm -p 80:80 $(IMAGE_NAME)-prod

.PHONY: build-dev build-prod run-dev run-prod restart-dev restart-prod
