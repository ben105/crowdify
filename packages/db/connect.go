package db

import (
	"log"

	"github.com/ben105/crowdify/packages/env"
	"github.com/gocql/gocql"
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

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create a cluster session: %v\n", err)
	}
	return &DbConnection{
		session,
	}
}
