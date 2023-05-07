#!/bin/sh

set -e

echo "run db migration"
echo $DB_SOURCE
echo $DB_DRIVER
echo $SERVER_ADDRESS
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"