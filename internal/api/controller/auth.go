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

func (c *AuthController) Register(ctx *gin.Context) {
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
	createdUser, httpErr := c.AuthUsecase.Register(context.Background(), user)
	if httpErr != nil {
		ctx.JSON(httpErr.Status(), httpErr)
		return
	}

	response := dto.RegisterResponse{User: createdUser}

	ctx.JSON(http.StatusCreated, response)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.NewHttpError(http.StatusBadRequest, errors.Validation, strings.Split(err.Error(), "\n")))
		return
	}

	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	tokenized, httpErr := c.AuthUsecase.Login(context.Background(), user)
	if httpErr != nil {
		ctx.JSON(httpErr.Status(), httpErr)
		return
	}

	response := dto.LoginResponse{
		User: tokenized,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) Me(ctx *gin.Context) {
	id := ctx.GetInt("X-User-ID")

	user, err := c.AuthUsecase.Me(context.Background(), id)
	if err != nil {
		ctx.JSON(err.Status(), err)
	}

	ctx.JSON(http.StatusOK, user)
}
