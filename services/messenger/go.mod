module github.com/ben105/crowdify/services/messenger

go 1.23.7

require (
	github.com/ben105/crowdify/packages/env v0.0.0-20250415195112-84b1c7874a2b
	github.com/ben105/crowdify/packages/message_queue v0.0.0-20250415195112-84b1c7874a2b
)

replace github.com/ben105/crowdify/packages/env => ../../packages/env

replace github.com/ben105/crowdify/packages/message_queue => ../../packages/message_queue

require (
	github.com/confluentinc/confluent-kafka-go/v2 v2.8.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
