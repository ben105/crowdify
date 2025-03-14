#!/bin/bash

set -e
set -o pipefail

retries=5
for i in $(seq 1 $retries); do
    if ! statuses=$(nodetool -u "$JMX_USER" -pw "$JMX_PASSWORD" -h "$CQLSH_HOST" status | grep -E "^[UD][NLRM]" | awk '{print $1}'); then
        echo "ERROR: Failed trying to run nodetool. Check the values for JMX_USER and JMX_PASSWORD"
        exit $!
    fi
    # Check if all statuses are "UN"
    if [[ $(echo "$statuses" | grep -v "UN") ]]; then
        echo "Current statuses:"
        nodetool status | grep -E "^[UD][NLRM]"
        sleep $i
    else
        echo "SUCCESS: All nodes are in UP/NORMAL state"
        echo "Node count: $(echo "$statuses" | wc -l)"
        break
    fi
done

if [[ "$i" -gt "$retries" ]]; then
  echo >&2 "ERROR: Failed to connect to cassandra at ${CQLSH_HOST}:${CQLSH_PORT}"
  exit 1
fi

for file in /scripts/*.cql; do
  [ -e "$file" ] || continue
  echo "Executing $file..."
  cqlsh --cqlversion "$CQLVERSION" -u "$USER" -p "$PASSWORD" -f "$file"
done

echo "Done."
