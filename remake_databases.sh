#!/bin/bash
docker-compose exec db mysql -u root --execute="DROP DATABASE dddsample;";
docker-compose exec db mysql -u root --execute="CREATE DATABASE dddsample;";