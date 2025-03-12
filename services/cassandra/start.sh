#!/bin/sh

docker network create cassandra 2> /dev/null
docker run --rm -d --name cassandra --hostname cassandra --network cassandra cassandra