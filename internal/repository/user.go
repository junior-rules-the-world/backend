package repository

import (
	"context"
	"events-organizator/internal/domain/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := "insert into users (display_name, username, email, password, team_id) values ($1, $2, $3, $4, $5) returning *;"

	err := user.BeforeCreate()
	if err != nil {
		return nil, err
	}

	tx := u.DB.MustBegin()
	dest := &models.User{}
	err = tx.QueryRowxContext(ctx, query, user.DisplayName, user.Username, user.Email, user.Password, user.TeamID).
		StructScan(dest)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (u *UserRepository) Update(ctx context.Context, id int, updated *models.User) error {
	query := "update users set display_name = coalesce(nullif($2, ''), display_name), role = coalesce(nullif($3, ''), role),team_id = $4,updated_at = now() where id = $1;"

	err := updated.BeforeUpdate()
	if err != nil {
		return err
	}
	dest := &models.User{}

	tx := u.DB.MustBegin()
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

func (u *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := "select * from users where username=$1;"
	dest := &models.User{}

	err := u.DB.QueryRowxContext(ctx, query, username).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "select * from users where email=$1"

	dest := &models.User{}

	err := u.DB.QueryRowxContext(ctx, query, email).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (u *UserRepository) FindById(ctx context.Context, id int) (*models.User, error) {
	query := "select * from users where id=$1"
	dest := &models.User{}

	err := u.DB.QueryRowxContext(ctx, query, id).StructScan(dest)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
