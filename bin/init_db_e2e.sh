#!/bin/bash
cd docker || exit
docker compose cp ./sql postgres:/var/lib/postgresql/data
docker compose exec postgres bash -c "psql -v ON_ERROR_STOP=1 --username \"$POSTGRES_USER\" --dbname postgres < /var/lib/postgresql/data/sql/e2e/reset_database.e2e.sql"
docker compose exec postgres bash -c "psql -v ON_ERROR_STOP=1 --username \"$POSTGRES_USER\" --dbname db_e2e < /var/lib/postgresql/data/sql/migrations/20220724145343_create_users_todos.up.sql"
