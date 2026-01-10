docker-db:
	docker-compose -f ./database/docker-compose.yml up -d
docker-db-down:
	docker-compose -f ./database/docker-compose.yml down -v

run-next:
	cd ./frontend && npm run dev
	