package jwt

import (
	"events-organizator/internal/domain/models"
	"events-organizator/internal/setup"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User, env *setup.Env) (string, error) {
	claims := &Claims{
		Username: user.Username,
		ID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 72)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(env.JWTSecret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
