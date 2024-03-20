package route

import (
	"events-organizator/internal/api/controller"
	"events-organizator/internal/repository"
	"events-organizator/internal/setup"
	"events-organizator/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewAuthRoute(env *setup.Env, db sqlx.DB, group *gin.RouterGroup) {
	r := repository.NewUserRepository(&db)
	c := controller.AuthController{
		AuthUsecase: usecase.NewAuthUsecase(r, *env),
		Env:         env,
	}

	group.POST("/register", c.Register)
	group.POST("/login", c.Login)
}
