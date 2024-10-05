#!/bin/bash

set -e

until psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "\q"; do
    >&2 echo "Postgres is unavailable - retrying"
    sleep 1
done

>&2 echo "Postgres is available - Begin initialization"

psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /docker-entrypoint-initdb.d/init-db.sql

echo "Initialization is successful!!"