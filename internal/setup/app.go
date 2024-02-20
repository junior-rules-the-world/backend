package setup

import (
	"events-organizator/internal/postgres"
)

type Application struct {
	Env      *Env
	Postgres *postgres.Postgres
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresConnection(app.Env)
	return *app
}
