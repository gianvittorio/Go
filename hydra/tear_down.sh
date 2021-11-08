#!/bin/bash

docker-compose down
docker volume rm $(docker volume ls --filter name=hydra)
unset POSTGRES_USER
unset POSTGRES_PASSWORD
unset POSTGRES_DB
unset POSTGRES_PORT
