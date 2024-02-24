package repository

import (
	"events-organizator/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db sqlx.DB
}

func (u *UserRepository) Create(user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Update(id int, updated *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FindByUsername(username string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FindById(username int) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}
