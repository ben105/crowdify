module github.com/ben105/crowdify/packages/kafka

go 1.23.7

require (
	github.com/ben105/crowdify/packages/env v0.0.0-20250318211338-c8fa4d5b79ec
	github.com/confluentinc/confluent-kafka-go/v2 v2.8.0
)

replace github.com/ben105/crowdify/packages/env => ../env

require github.com/joho/godotenv v1.5.1 // indirect
