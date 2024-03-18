package setup

import (
	"events-organizator/internal/postgres"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func NewPostgresConnection(env *Env) *postgres.Postgres {
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	connection, err := pgx.ParseConfig(postgresURL)
	if err != nil {
		log.Fatalf("Error while parsing postgresUrl: %s", err)
	}
	conn, err := postgres.NewConn(connection.ConnString())
	if err != nil {
		log.Fatalf("Error while connecting to conn: %s", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Got an error while connecting to conn: %s", err)
	}

	return conn
}
