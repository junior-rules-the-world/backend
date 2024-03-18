package main

import (
	"events-organizator/internal/api/route"
	"events-organizator/internal/setup"
	gin2 "github.com/gin-gonic/gin"
	"time"
)

func main() {
	app := setup.App()

	env := app.Env

	gin := gin2.Default()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	route.Setup(env, timeout, *app.Postgres.Connection, gin)

	gin.Run(env.Address)
}
