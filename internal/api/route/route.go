package route

import (
	"events-organizator/internal/api/middleware"
	"events-organizator/internal/setup"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(env *setup.Env, db sqlx.DB, gin *gin.Engine) {
	publicGroup := gin.Group("")

	NewAuthRoute(env, db, publicGroup)

	protectedGroup := gin.Group("")
	protectedGroup.Use(middleware.Jwt(env))
}
