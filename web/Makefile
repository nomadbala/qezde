local-dev:
	bun run dev

start-dev:
	docker-compose -f docker-compose.dev.yml up

down-dev:
	docker-compose -f docker-compose.dev.yml down

restart-dev: down-dev start-dev

start-prod:
	docker-compose -f docker-compose.prod.yml up

down-prod:
	docker-compose -f docker-compose.prod.yml down

restart-prod:
	make down-prod
	make start-prod

.PHONY: local-dev start-dev down-dev start-prod down-prod restart-dev restart-prod