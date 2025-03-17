# Running Locally

## Starting Services (Emulator)

To emulate services required for running Crowdify locally, run the `make em` command. Among other things, this will boot up the Cassandra database running the necessary bootstrap scripts and all schema migrations.

## Tools

### Cassandra Tools

You can add fake tracks to the Cassandra database. At the moment this doesn't do much, except prove to yourself that the database is actually running with the right tables.

```sh
cd tools/cli
go run . track add "We Will Rock You"
```

If you want to see the tracks in the table, you can connect to the database and run a select.

```sh
cd tools/cassandra
./connect.sh

cqlsh> SELECT * FROM crowdify.unprocessed_tracks;
```