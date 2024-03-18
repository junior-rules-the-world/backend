package route

import (
	"events-organizator/internal/setup"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func Setup(env *setup.Env, timeout time.Duration, db sqlx.DB, gin *gin.Engine) {
	publicGroup := gin.Group("")

	NewAuthRoute(env, timeout, db, publicGroup)
}
