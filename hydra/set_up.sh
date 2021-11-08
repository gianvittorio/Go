#!/bin/bash

export POSTGRES_USER=test
export POSTGRES_PASSWORD=test
export POSTGRES_DB=test
export POSTGRES_PORT=5432

docker-compose up -d --build

sleep 10

docker-compose exec -T hydra-postgresql-db psql -d ${POSTGRES_DB} -U ${POSTGRES_USER} -p ${POSTGRES_PORT} << EOF
    create table if not exists personnel (
        id serial primary key, 
        name varchar(45) not null,
        security_clearance int not null,
        position varchar(45) not null
    )
EOF
