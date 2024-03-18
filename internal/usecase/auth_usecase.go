package usecase

import (
	"context"
	"events-organizator/internal/domain/models"
	"events-organizator/pkg/errors"
	"net/http"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}

type AuthUsecase struct {
	repo       UserRepository
	ctxTimeout time.Duration
}

func NewAuthUsecase(repo UserRepository, timeout time.Duration) *AuthUsecase {
	return &AuthUsecase{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

func (uc *AuthUsecase) Register(ctx context.Context, user *models.User) (*models.User, *errors.HttpError) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
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
