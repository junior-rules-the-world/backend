package middleware

import (
	"events-organizator/internal/setup"
	"events-organizator/pkg/errors"
	"events-organizator/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Jwt(env *setup.Env) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		bearer := strings.Split(header, " ")
		if len(bearer) == 2 {
			token := bearer[1]

			authorized, err := jwt.IsValid(token, env)

			if !authorized {
				ctx.JSON(http.StatusUnauthorized, errors.NewHttpError(http.StatusUnauthorized, errors.Unauthorized, err.Error()))
				ctx.Abort()
				return
			}

			id, err := jwt.ExtractUserID(token, env)

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, errors.NewHttpError(http.StatusUnauthorized, errors.Unauthorized, err.Error()))
				ctx.Abort()
				return
			}

			ctx.Set("X-User-ID", id)
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, errors.NewHttpError(http.StatusUnauthorized, errors.Unauthorized, nil))
		ctx.Abort()
	}
}
