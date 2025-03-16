package db

import (
	"log"
	"os"

	"github.com/ben105/crowdify/packages/env"
	"github.com/gocql/gocql"
)

func Connect() *gocql.Session {
	cluster := gocql.NewCluster(env.CassandraHost)
	cluster.Keyspace = "crowdify"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create a cluster session: %v\n", err)
	}
	return session
}
