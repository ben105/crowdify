#!/bin/bash

set -e

TIMEOUT=10
KEYSPACE_NAME="crowdify"

# if ! cqlsh --cqlversion "${CQLVERSION:-3.4.7}" -u "${CASSANDRA_USERNAME:-cassandra}" -p "${CASSANDRA_PASSWORD:-cassandra}" --connect-timeout=$TIMEOUT -e "SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = '$KEYSPACE_NAME';" >/dev/null 2>&1; then
if ! cqlsh --cqlversion "${CQLVERSION:-3.4.7}" -u "${CASSANDRA_USERNAME:-cassandra}" -p "${CASSANDRA_PASSWORD:-cassandra}" --connect-timeout=$TIMEOUT -e "SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = '$KEYSPACE_NAME';"; then
    echo "Healthcheck failed:  Keyspace '$KEYSPACE_NAME' does not exist"
    exit 1
fi

echo "Healthcheck passed: Keyspace '$KEYSPACE_NAME' exists"