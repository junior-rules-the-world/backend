package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Connection *sqlx.DB
}

func NewConn(connection string) (*Postgres, error) {
	conn, err := sqlx.Connect("pgx", connection)

	pg := Postgres{
		Connection: conn,
	}

	return &pg, err
}

func (p *Postgres) Ping() error {
	return p.Connection.Ping()
}
