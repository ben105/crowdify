# Cassandra

## Running Things Locally

Spinning up a Cassandra service locally requires starting the cassandra container, followed by runing some scripts. Spin up the server first by running:

```sh
./start.sh
```

It will take a few minutes for the Cassandra service to start and be healthy. To properly check health status we use `nodetool` in our `docker/cql_runner` Docker image. This uses the JMX service on port 7199 to check that the service is both "up" and "normal" before running CQL commands. This is abstracted away in the `run_scripts.sh` bash script.

The Cassandra service is created with a default role `cassandra`, with the default username/password of cassandra/cassandra. Add yourself as a superuser by modifying the `bootstrap_scripts/bootstrap.cql` and replacing the relevant values.

You cannot drop the default `cassandra` role while logged in as this very same role. So the steps are split into two steps:
1. Run the bootstrap scripts: `./run_scripts.sh bootstrap`
2. Run the regular scripts as your new superuser role: `USER=youruser PASSWORD=yourpassword ./run_scripts.sh`

## Java Management Extensions (JMX)

The local Cassandra service is configured with LOCAL_JMX=no. This enables remote JMX connections. The Cassandra container doesn't publish any ports to the host, but this is still required for communication between containers in the cassandra network.

The default user/password values are admin/admin. The `start.sh` script uploads the `jmxremote.password` file to the Cassandra server with these values, but they can be overriden by the `$JMX_USER` and `$JMX_PASSWORD` environment variables respectively. Sane values should be provided in production.