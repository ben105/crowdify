FROM cassandra:latest AS cassandra-base

COPY cassandra.yaml /etc/cassandra/

COPY scripts /scripts

FROM migrate/migrate:latest AS migration

FROM cassandra-base

COPY migrations /migrations
COPY --from=migration /usr/local/bin/migrate /usr/local/bin/migrate

COPY healthcheck.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/healthcheck.sh

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER nobody

HEALTHCHECK --interval=10s --timeout=10s --start-period=30s --retries=10 \
  CMD /usr/local/bin/healthcheck.sh

ENTRYPOINT ["/entrypoint.sh"]