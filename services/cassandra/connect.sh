#!/bin/sh

source ./env.sh

docker run \
  --rm \
  -it \
  --network cassandra \
  nuvo/docker-cqlsh \
  cqlsh -u "$USER" -p "$PASSWORD" "$CASSANDRA_HOST" "$CASSANDRA_PORT" --cqlversion="$CQLVERSION"