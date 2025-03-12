#!/bin/sh

source ./version.sh

docker run --rm -it --network cassandra nuvo/docker-cqlsh cqlsh cassandra 9042 --cqlversion=$CQLVERSION