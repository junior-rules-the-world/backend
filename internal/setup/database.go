package setup

import (
	"context"
	"events-organizator/internal/postgres"
	"fmt"
	"log"
	"time"
)

func NewPostgresConnection(env *Env) *postgres.Postgres {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postgresURL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", env.DBUser, env.DBPassword, env.DBPort, env.DBName)
	postgres, err := postgres.NewConn(postgresURL)
	if err != nil {
		log.Fatalf("Error while connecting to postgres: %s", err)
	}

	err = postgres.Ping(ctx)
	if err != nil {
		log.Fatalf("Got an error while connecting to postgres: %s", err)
	}

	return postgres
}
