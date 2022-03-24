main:
	docker-compose exec go go run api/main.go

remake:
	docker-compose exec db mysql -u root --execute="DROP DATABASE dddsample;";
	docker-compose exec db mysql -u root --execute="CREATE DATABASE dddsample;";

migrate:
	docker-compose exec go sql-migrate up /
	docker-compose exec go go run database/seeds/seed.go

go:
	docker-compose exec go bash

test:
	docker-compose exec go go clean -testcache /
	docker-compose exec go gotest -v ./...