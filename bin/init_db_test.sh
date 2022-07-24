#!/bin/bash
cd docker || exit
docker compose cp ./sql postgres:/var/lib/postgresql/data
docker compose exec postgres bash -c "psql -v ON_ERROR_STOP=1 --username \"$POSTGRES_USER\" --dbname postgres < /var/lib/postgresql/data/sql/test/reset_database.test.sql"
docker compose exec postgres bash -c "psql -v ON_ERROR_STOP=1 --username \"$POSTGRES_USER\" --dbname db_test < /var/lib/postgresql/data/sql/migrations/20220724145343_create_users_todos.up.sql"
