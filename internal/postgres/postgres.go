package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Client interface {
	Connect(ctx context.Context) error
	Ping(ctx context.Context) error
}

type Postgres struct {
	Connection *pgx.Conn
}

func NewConn(connection string) (*Postgres, error) {
	conn, err := pgx.Connect(context.Background(), connection)

	return &Postgres{Connection: conn}, err
}

func (client Postgres) Ping(ctx context.Context) error {
	return client.Connection.Ping(ctx)
}
