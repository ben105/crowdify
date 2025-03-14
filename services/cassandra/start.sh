#!/bin/sh

source ./env.sh

docker network create cassandra 2> /dev/null
docker run \
  --rm \
  -d \
  -e LOCAL_JMX=no \
  -e JMX_HOSTNAME="$CASSANDRA_HOST" \
  -v "$(pwd)/cassandra.yaml:/etc/cassandra/cassandra.yaml" \
  --name cassandra \
  --hostname "$CASSANDRA_HOST" \
  --network cassandra \
  cassandra:"$CASSANDRA_VERSION"

container_id=$(docker ps -aqf "name=cassandra")

cat <<EOL >jmxremote.password
$JMX_USER $JMX_PASSWORD
EOL
docker cp jmxremote.password ${container_id}:/etc/cassandra/jmxremote.password
rm jmxremote.password 
