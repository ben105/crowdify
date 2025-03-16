# Kafka Tool

## Simple Usage

Simple tool for testing Kafka deployment. Usage:

```sh
go run . consume
```

```sh
go run . publish "Any message here"
```

## Environment Variables

You can update the Kafka broker, topic, group, etc by updating the values in your .env file at the root of htis project. For example:

```sh
TOPIC=local-topic
BROKER=localhost:9092
GROUP_ID=local-consumer-group
GROUP_INSTANCE_ID=local-consumer-group-instance
```

Default values will be used if you omit any of the above. Check out the env package in `packages/env` to see default values.