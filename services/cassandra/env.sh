#!/bin/sh

# If there's ever skew with the CQL version, you can figure out which version to use by running this command:
# docker run -it --network cassandra --rm cassandra cqlsh -e "show version" cassandra | awk '{ print $9 }' | head -1
export CQLVERSION=3.4.7

export CASSANDRA_VERSION=5.0.3
export CASSANDRA_HOST=${CASSANDRA_HOST:-cassandra}
export CASSANDRA_PORT=${CASSANDRA_PORT:-9042}

export JMX_USER=${JMX_USER:-admin}
export JMX_PASSWORD=${JMX_PASSWORD:-admin}
