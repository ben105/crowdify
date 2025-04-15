module github.com/ben105/crowdify/packages/message_queue

go 1.23.7

require (
	github.com/ben105/crowdify/packages/env v0.0.0-20250318211338-c8fa4d5b79ec
	github.com/confluentinc/confluent-kafka-go/v2 v2.8.0
	github.com/stretchr/testify v1.9.0
)

replace github.com/ben105/crowdify/packages/env => ../env

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
