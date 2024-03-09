package usecase

import (
	"context"
	"events-organizator/internal/domain"
	"events-organizator/pkg/errors"
	"net/http"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, id int, updated *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindById(ctx context.Context, id int) (*domain.User, error)
}

type UserUsecase struct {
	repo       UserRepository
	ctxTimeout time.Duration
}

func NewUserUsecase(repo UserRepository, timeout time.Duration) *UserUsecase {
	return &UserUsecase{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

func (uc *UserUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
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

	if err = user.BeforeCreate(); err != nil {
		return nil, errors.NewHttpError(http.StatusInternalServerError, errors.ServerError, err)
	}

	user, err = uc.repo.Create(ctx, user)
	if err != nil {
		return nil, errors.NewHttpError(http.StatusInternalServerError, errors.ServerError, err)
	}

	return user, nil
}
