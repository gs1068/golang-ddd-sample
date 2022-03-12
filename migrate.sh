docker-compose exec go sql-migrate up /
docker-compose exec go go run database/seeds/seed.go

test