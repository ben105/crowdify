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

HEALTHCHECK --interval=10s --timeout=10s --start-period=150s --retries=30 \
  CMD /usr/local/bin/healthcheck.sh

USER nobody

ENTRYPOINT ["/entrypoint.sh"]