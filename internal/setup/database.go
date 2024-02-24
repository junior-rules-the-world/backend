package setup

import (
	"events-organizator/internal/postgres"
	"fmt"
	"log"
)

func NewPostgresConnection(env *Env) *postgres.Postgres {
	postgresURL := fmt.Sprintf("conn://%s:%s@%s:%s/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	conn, err := postgres.NewConn(postgresURL)
	if err != nil {
		log.Fatalf("Error while connecting to conn: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Got an error while connecting to conn: %s", err)
	}

	return conn
}
