package main

import (
	"events-organizator/internal/setup"
	gin2 "github.com/gin-gonic/gin"
)

func main() {
	app := setup.App()

	env := app.Env

	gin := gin2.Default()

	gin.Run(env.Address)
}
