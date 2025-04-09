package db

import (
	"fmt"
	"log"
	"time"

	"github.com/ben105/crowdify/packages/env"
	"github.com/gocql/gocql"
)

var (
	maxRetries = 5
	retryDelay = 2 * time.Second
)

type DbConnection struct {
	session *gocql.Session
}

func Connect() *DbConnection {
	cluster := gocql.NewCluster(env.GetCassandraHost())
	cluster.Keyspace = "crowdify"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: env.GetCassandraUsername(),
		Password: env.GetCassandraPassword(),
	}

	var session *gocql.Session
	var err error

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempt %d to connect to Cassandra...\n", i+1)
		session, err = cluster.CreateSession()
		if err == nil {
			log.Println("Successfully connected to Cassandra.")
			return &DbConnection{
				session,
			}
		}

		log.Printf("Failed to connect to Cassandra (attempt %d/%d): %v\n", i+1, maxRetries, err)

		if i < maxRetries-1 {
			log.Printf("Retrying in %v...\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	panic(fmt.Sprintf("Failed to connect to Cassandra after %d attempts: %v\n", maxRetries, err))
}
