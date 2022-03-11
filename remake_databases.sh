docker-compose exec db mysql -u root -e "DROP DATABASE dddsample;";
docker-compose exec db mysql -u root -e "CREATE DATABASE dddsample;";