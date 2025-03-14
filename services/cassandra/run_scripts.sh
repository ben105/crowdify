#!/bin/sh

source ./env.sh

SCRIPTS_SRC="$(pwd)/scripts"

if [[ $1 == "bootstrap" ]]; then
  USER=${USER:-cassandra}
  PASSWORD=${PASSWORD:-cassandra}
  SCRIPTS_SRC="$(pwd)/bootstrap_scripts"
fi

pushd docker/cql_runner
docker build \
  --build-arg CASSANDRA_VERSION="$CASSANDRA_VERSION" \
  --build-arg CQLVERSION="$CQLVERSION" \
  -t cql-runner \
  .
popd

docker run \
  --rm \
  --network cassandra \
  -v "${SCRIPTS_SRC}:/scripts" \
  -e USER="$USER" \
  -e PASSWORD="$PASSWORD" \
  -e CQLSH_HOST="$CASSANDRA_HOST" \
  -e CQLSH_PORT="$CASSANDRA_PORT" \
  -e CQLVERSION="$CQLVERSION" \
  -e JMX_USER="$JMX_USER" \
  -e JMX_PASSWORD="$JMX_PASSWORD" \
  cql-runner
