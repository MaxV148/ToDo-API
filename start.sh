#!/bin/sh

set -e

echo "run db migration"
echo "$DB_SOURCE"
/app/migrate -path /app/migration -database "$DB_SOURCE_PROD" -verbose up

echo "start the api"
exec "$@"