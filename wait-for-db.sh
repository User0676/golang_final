#!/bin/sh

# Wait for PostgreSQL to be ready
while ! nc -z db 5432; do
  echo "Waiting for the database..."
  sleep 1
done

echo "Database is ready"
exec "$@"
