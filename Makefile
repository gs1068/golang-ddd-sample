docker:
	docker-compose up -d --build

main:
	docker-compose exec go fresh -c fresh.conf

remake:
	docker-compose exec db mysql -u root --execute="DROP DATABASE dddsample;";
	docker-compose exec db mysql -u root --execute="CREATE DATABASE dddsample;";

migrate:
	docker-compose exec go sql-migrate up /
	docker-compose exec go go run database/seeds/seed.go