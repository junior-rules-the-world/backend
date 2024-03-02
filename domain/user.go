package domain

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID          int       `json:"id" db:"user_id"`
	DisplayName *string   `json:"display_name" db:"display_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password"`
	Role        *string   `json:"role,omitempty"`
	TeamID      *int      `json:"team_id,omitempty" db:"team_id"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type TokenizedUser struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type UserRepository interface {
	Create(user *User) (User, error)
	Update(id int, updated *User) error
	FindByUsername(username string) (User, error)
	FindById(username int) (User, error)
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	*u.DisplayName = strings.ToLower(strings.TrimSpace(*u.DisplayName))
	u.Password = strings.ToLower(strings.TrimSpace(u.Password))

	if u.DisplayName == nil {
		*u.DisplayName = u.Username
	}

	if err := u.HashPassword(); err != nil {
		return err
	}

	return nil
}
func (u *User) BeforeUpdate() error {
	*u.DisplayName = strings.ToLower(strings.TrimSpace(*u.DisplayName))

	if u.DisplayName == nil {
		*u.DisplayName = u.Username
	}

	return nil
}
