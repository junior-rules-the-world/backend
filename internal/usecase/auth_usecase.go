package usecase

import (
	"context"
	"events-organizator/internal/domain/models"
	"events-organizator/internal/setup"
	"events-organizator/pkg/errors"
	"events-organizator/pkg/jwt"
	"net/http"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindById(ctx context.Context, id int) (*models.User, error)
}

type AuthUsecase struct {
	repo    UserRepository
	env     setup.Env
	timeout time.Duration
}

func NewAuthUsecase(repo UserRepository, env setup.Env) *AuthUsecase {
	return &AuthUsecase{
		repo:    repo,
		env:     env,
		timeout: time.Duration(env.ContextTimeout) * time.Second,
	}
}

func (uc *AuthUsecase) Register(ctx context.Context, user *models.User) (*models.User, *errors.HttpError) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	userExists, err := uc.repo.FindByUsername(ctx, user.Username)
	if userExists != nil || err == nil {
		return nil, errors.NewHttpError(http.StatusBadRequest, errors.UsernameAlreadyUsed, nil)
	}
	userExists, err = uc.repo.FindByEmail(ctx, user.Email)
	if userExists != nil || err == nil {
		return nil, errors.NewHttpError(http.StatusBadRequest, errors.EmailAlreadyUsed, nil)
	}

	user, err = uc.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusInternalServerError, errors.ServerError, err)
	}

	return user, nil
}

func (uc *AuthUsecase) Login(ctx context.Context, u *models.User) (*models.TokenizedUser, *errors.HttpError) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	// Finding user by all possible ways
	user, err := uc.repo.FindByUsername(ctx, u.Username)
	if user == nil || err != nil {
		return nil, errors.NewHttpError(http.StatusNotFound, errors.NotFound, nil)
	}

	if err = user.ComparePasswords(u.Password); err != nil {
		return nil, errors.NewHttpError(http.StatusUnauthorized, errors.Unauthorized, err.Error())
	}

	token, err := jwt.GenerateJWT(user, &uc.env)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusInternalServerError, errors.ServerError, nil)
	}

	return &models.TokenizedUser{
		User:  user,
		Token: token,
	}, nil
}

func (uc *AuthUsecase) Me(ctx context.Context, id int) (*models.User, errors.HttpErr) {
	ctx, cancel := context.WithTimeout(ctx, uc.timeout)
	defer cancel()

	user, err := uc.repo.FindById(ctx, id)

	if err != nil {
		return nil, errors.NewHttpError(http.StatusNotFound, errors.NotFound, nil)
	}

	return user, nil
}
