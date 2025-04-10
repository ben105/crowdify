module github.com/ben105/crowdify/services/crowdify

go 1.23.7

require (
	github.com/ben105/crowdify/packages/db v0.0.0-20250318211338-c8fa4d5b79ec
	github.com/ben105/crowdify/packages/kafka v0.0.0-00010101000000-000000000000

)

require (
	github.com/ben105/crowdify/packages/env v0.0.0-20250318211338-c8fa4d5b79ec // indirect
	github.com/confluentinc/confluent-kafka-go/v2 v2.8.0 // indirect
	github.com/gocql/gocql v1.7.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/ben105/crowdify/packages/db => ../../packages/db

replace github.com/ben105/crowdify/packages/env => ../../packages/env

replace github.com/ben105/crowdify/packages/kafka => ../../packages/kafka
