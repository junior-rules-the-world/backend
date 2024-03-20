package main

import (
	"events-organizator/internal/api/route"
	"events-organizator/internal/setup"
	gin2 "github.com/gin-gonic/gin"
)

func main() {
	app := setup.App()

	env := app.Env

	gin := gin2.Default()

	route.Setup(env, *app.Postgres.Connection, gin)

	gin.Run(env.Address)
}
