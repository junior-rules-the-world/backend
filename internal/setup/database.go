package setup

import (
	"events-organizator/pkg/postgres"
	"fmt"
	"log"
)

func NewPostgresConnection(env *Env) *postgres.Postgres {
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	postgres, err := postgres.NewConn(postgresURL)
	if err != nil {
		log.Fatalf("Error while connecting to postgres: %s", err)
	}

	err = postgres.Ping()
	if err != nil {
		log.Fatalf("Got an error while connecting to postgres: %s", err)
	}

	return postgres
}
