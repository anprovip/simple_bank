#!/bin/sh

echo "waiting for db..."
/app/wait-for.sh db:5432

echo "run db migration..."
/app/migrate -path ./db/migration -database "$DB_SOURCE" up

echo "start app..."
exec "$@"