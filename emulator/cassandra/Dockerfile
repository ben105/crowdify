FROM cassandra:latest AS cassandra-base

COPY cassandra.yaml /etc/cassandra/

COPY scripts /scripts

FROM migrate/migrate:latest AS migration

FROM cassandra-base

COPY migrations /migrations
COPY --from=migration /usr/local/bin/migrate /usr/local/bin/migrate

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER nobody

ENTRYPOINT ["/entrypoint.sh"]