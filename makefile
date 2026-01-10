include backend/.env
export

APP_NAME=live_music
# DB_URL=postgres://postgres:postgres@localhost:5432/live_music?sslmode=disable
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS=./backend/db/migrations
DOCKER_COMPOSE=database/postgres/docker-compose.yml


.PHONY: run db-up db-down migrate-up migrate-down reset

db-up:
	docker-compose -f $(DOCKER_COMPOSE) up -d
db-down:
	docker-compose -f $(DOCKER_COMPOSE) down -v

run-next:
	cd ./frontend && npm run dev
run-go:
	cd ./backend && go run main.go


## Run migrations
migrate-up:
	migrate -path $(MIGRATIONS) -database "$(DB_URL)" up

## Rollback last migration
migrate-down:
	migrate -path $(MIGRATIONS) -database "$(DB_URL)" down 

## Full reset (dangers but useful)
reset:
	docker compose -f $(DOCKER_COMPOSE) down -v
	rm -rf database/postgres/data
