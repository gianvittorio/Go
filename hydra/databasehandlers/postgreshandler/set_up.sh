#!/bin/bash

export POSTGRES_USER=test
export POSTGRES_PASSWORD=test
export POSTGRES_DB=test
export POSTGRES_PORT=5432

docker-compose up -d --build

sleep 10

docker-compose exec -T hydra-postgresql-db psql -d ${POSTGRES_DB} -U ${POSTGRES_USER} -p ${POSTGRES_PORT} << EOF
    create sequence if not exists personnel_id_seq start 1;
    
    create table if not exists personnel (
        id int not null default nextval('personnel_id_seq')::regclass, 
        name varchar(45) not null,
        security_clearance int not null,
        position varchar(45) not null,
        primary key (id)
    );

    insert into personnel (name, security_clearance, position) values('Isis Adcox', 5, 'Engineer I');
    insert into personnel (name, security_clearance, position) values('You Chaloux', 8, 'Engineer II');
    insert into personnel (name, security_clearance, position) values('Lorette Gee', 2, 'Assistant Pilot');
    insert into personnel (name, security_clearance, position) values('Telma Rosas', 5, 'Mechanic');
    insert into personnel (name, security_clearance, position) values('Lynsey Christman', 3, 'Technician');
EOF
