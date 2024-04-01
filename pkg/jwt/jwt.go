package jwt

import (
	"events-organizator/internal/domain/models"
	"events-organizator/internal/setup"
	"fmt"
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

	tokenString, err := token.SignedString([]byte(env.JWTSecret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsValid(token string, env *setup.Env) (bool, error) {
	_, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %s", tk.Header["alg"])
		}

		return []byte(env.JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractUserID(token string, env *setup.Env) (int, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %s", tk.Header["alg"])
		}

		return []byte(env.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := t.Claims.(*Claims)

	if !ok && !t.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}
