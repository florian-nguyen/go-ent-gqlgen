#!/bin/bash
cat docker/sql/seed/01_init.sql | docker exec -i postgres-db psql -b db