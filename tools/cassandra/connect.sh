#!/bin/sh

# TODO: Replace with Go tool? May be more convenient to allow passing these values in as arguments.

source ../../.env

docker run \
  --rm \
  -it \
  --network cassandra \
  nuvo/docker-cqlsh \
  cqlsh -u "${CASSANDRA_USERNAME:-cassandra}" -p "${CASSANDRA_PASSWORD:-cassandra}" --cqlversion="${CQLVERSION:-3.4.7}" \
    "${CASSANDRA_HOST:-cassandra}" \
    "${CASSANDRA_PORT:-9042}"