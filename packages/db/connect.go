package db

import (
	"log"

	"github.com/ben105/crowdify/packages/env"
	"github.com/gocql/gocql"
)

func Connect() *gocql.Session {
	cluster := gocql.NewCluster(env.GetCassandraHost())
	cluster.Keyspace = "crowdify"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: env.GetCassandraUsername(),
		Password: env.GetCassandraPassword(),
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create a cluster session: %v\n", err)
	}
	return session
}
