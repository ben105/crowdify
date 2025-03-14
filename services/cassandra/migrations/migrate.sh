docker run -v "$(pwd):/migrations" --network cassandra migrate/migrate \
    -path=/migrations/ -database cassandra://cassandra/store up