package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)
	err := godotenv.Load(filepath.Join(moduleDir, "../../.env"))
	if err != nil {
		fmt.Println("Could not find an .env file. Only defaults will be used.")
	}
}

func GetTopic() string {
	return GetEnv("TOPIC", "local-topic")
}

func GetBroker() string {
	return GetEnv("BROKER", "localhost:9092")
}

func GetDeadLetterQueueTopic() string {
	return GetEnv("DEAD_LETTER_QUEUE_TOPIC", "local-dead-letter-queue-topic")
}

func GetGroupId() string {
	return GetEnv("GROUP_ID", "local-consumer-group")
}

func GetGroupInstanceId() string {
	return GetEnv("GROUP_INSTANCE_ID", "local-consumer-group-instance")
}

func GetCqlVersion() string {
	return GetEnv("CQLVERSION", "3.4.7")
}

func GetCassandraHost() string {
	return GetEnv("CASSANDRA_HOST", "cassandra")
}

func GetCassandraPort() string {
	return GetEnv("CASSANDRA_PORT", "9042")
}

func GetCassandraUsername() string {
	return GetEnv("CASSANDRA_USERNAME", "cassandra")
}

func GetCassandraPassword() string {
	return GetEnv("CASSANDRA_PASSWORD", "cassandra")
}
