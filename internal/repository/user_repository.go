package repository

import (
	"context"
	"events-organizator/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		*db,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := "insert into users (display_name, username, email, password, team_id) values ($1, $2, $3, $4, $5);"

	err := user.BeforeCreate()
	if err != nil {
		return nil, err
	}

	dest := domain.User{}

	tx := u.db.MustBegin()
	err = tx.
		QueryRowxContext(ctx, query, user.DisplayName, user.Username, user.Password, user.TeamID).
		StructScan(u)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (u *UserRepository) Update(ctx context.Context, id int, updated *domain.User) error {
	query := "update users set display_name = coalesce(nullif($2, ''), display_name), role = coalesce(nullif($3, ''), role),team_id = $4,updated_at = now() where id = $1;"

	err := updated.BeforeUpdate()
	if err != nil {
		return err
	}
	dest := domain.User{}

	tx := u.db.MustBegin()
	err = tx.QueryRowxContext(ctx, query, id, updated.DisplayName, updated.Role, updated.TeamID).StructScan(dest)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := "select from users where username like $1;"
	dest := domain.User{}

	err := u.db.QueryRowxContext(ctx, query, username).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := "select from users where email=$1"

	dest := domain.User{}

	err := u.db.QueryRowxContext(ctx, query, email).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (u *UserRepository) FindById(ctx context.Context, id int) (*domain.User, error) {
	query := "select from users where id=$1"
	dest := domain.User{}

	err := u.db.QueryRowxContext(ctx, query, id).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}