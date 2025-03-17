#!/bin/bash

TIMEOUT=10
KEYSPACE_NAME="crowdify"

# Check if the keyspace exists
echo "Checking if keyspace '$KEYSPACE_NAME' exists..."
KEYSPACE_CHECK=$(cqlsh $CASSANDRA_HOST $CASSANDRA_PORT -u "$CASSANDRA_USERNAME" -p "$CASSANDRA_PASSWORD" -e "SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = '$KEYSPACE_NAME';" --connect-timeout=$TIMEOUT)

# Parse the output to determine if the keyspace exists
if echo "$KEYSPACE_CHECK" | grep -q "$KEYSPACE_NAME"; then
  echo "Healthcheck passed: Keyspace '$KEYSPACE_NAME' exists"
  exit 0
else
  echo "Healthcheck failed: Keyspace '$KEYSPACE_NAME' does not exist"
  exit 1
fi