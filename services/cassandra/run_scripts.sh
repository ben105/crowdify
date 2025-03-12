#!/bin/sh

CQLVERSION=`docker run -it --network cassandra --rm cassandra cqlsh -e "show version" cassandra | awk '{ print $9 }' | head -1`

docker run --rm --network cassandra -v "$(pwd)/scripts:/scripts" -e CQLSH_HOST=cassandra -e CQLSH_PORT=9042 -e CQLVERSION=$CQLVERSION nuvo/docker-cqlsh
