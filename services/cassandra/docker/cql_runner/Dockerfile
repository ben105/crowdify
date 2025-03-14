ARG CASSANDRA_VERSION=latest
FROM cassandra:$CASSANDRA_VERSION

WORKDIR /scripts

ARG CQLVERSION
ENV CQLVERSION=${CQLVERSION}

RUN if [ -z "$CQLVERSION" ]; then \
        echo "Error: CQLVERSION is not set!"; \
        exit 1; \
    fi

ENV CQLSH_HOST="cassandra" \
    CQLSH_PORT="9042" \
    CASSANDRA_VERSION="5.0.3"

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
