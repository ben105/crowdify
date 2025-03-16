package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Topic,
	Broker,
	GroupId,
	GroupInstanceId,
	CqlVersion,
	CassandraHost,
	CassandraPort string
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Could not find an .env file. Only defaults will be used.")
	}

	// Kafka related environment variables
	Topic = GetEnv("TOPIC", "local-topic")
	Broker = GetEnv("BROKER", "localhost:9092")
	GroupId = GetEnv("GROUP_ID", "local-consumer-group")
	GroupInstanceId = GetEnv("GROUP_INSTANCE_ID", "local-consumer-group-instance")

	// Cassandra related environment variables
	CqlVersion = GetEnv("CQLVERSION", "3.4.7")
	CassandraHost = GetEnv("CASSANDRA_HOST", "cassandra")
	CassandraPort = GetEnv("CASSANDRA_PORT", "9042")
}
