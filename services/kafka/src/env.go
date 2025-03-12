package main

import "os"

var (
	topic           = getEnv("TOPIC", "local-topic")
	broker          = getEnv("BROKER", "localhost:9092")
	groupId         = getEnv("GROUP_ID", "local-consumer-group")
	groupInstanceId = getEnv("GROUP_INSTANCE_ID", "local-consumer-group-instance")
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
