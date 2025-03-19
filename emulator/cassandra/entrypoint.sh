#!/bin/bash

set -e
set -o pipefail

cassandra -f &
CASSANDRA_PID=$!

retry_count=0
max_retries=10
until cqlsh --cqlversion "$CQLVERSION" -u "$CASSANDRA_USERNAME" -p "$CASSANDRA_PASSWORD" -e "describe keyspaces" >/dev/null 2>&1; do
  retry_count=$((retry_count+1))
  if [ "$retry_count" -gt "$max_retries" ]; then
    echo "Failed to start Cassandra"
    exit 1
  fi
  echo "Waiting for Cassandra to start..."
  sleep 10
done

echo "Cassandra is up and running"

echo "Running admin scripts..."
for file in /scripts/*.cql; do
  [ -e "$file" ] || continue
  echo "Executing $file..."
  cqlsh --cqlversion "$CQLVERSION" -u "$CASSANDRA_USERNAME" -p "$CASSANDRA_PASSWORD" -f "$file"
done


echo "Running migrations..."
/usr/local/bin/migrate -path=/migrations/ -database "cassandra://cassandra/crowdify?username=${CASSANDRA_USERNAME}&password=${CASSANDRA_PASSWORD}&x-multi-statement=true" up

if kill -0 $CASSANDRA_PID 2>/dev/null; then
  echo "Migrations completed successfully"
  wait $CASSANDRA_PID
else
  echo "Cassandra process died during migration"
  exit 1
fi
