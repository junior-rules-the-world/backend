package controller

import (
	"context"
	"events-organizator/internal/domain/dto"
	"events-organizator/internal/domain/models"
	"events-organizator/internal/setup"
	"events-organizator/internal/usecase"
	"events-organizator/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthController struct {
	AuthUsecase *usecase.AuthUsecase
	Env         *setup.Env
}

func (u *AuthController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.NewHttpError(http.StatusBadRequest, errors.Validation, strings.Split(err.Error(), "\n")))
		return
	}

	user := &models.User{
		DisplayName: request.DisplayName,
		Username:    request.Username,
		Email:       request.Email,
		Password:    request.Password,
		Role:        request.Role,
		TeamID:      request.TeamID,
	}
	createdUser, httpErr := u.AuthUsecase.Register(context.Background(), user)
	if httpErr != nil {
		ctx.JSON(httpErr.Status(), httpErr)
		return
	}

	response := dto.RegisterResponse{User: createdUser}

	ctx.JSON(http.StatusCreated, response)
}
