package dto

import "events-organizator/internal/domain/models"

type RegisterRequest struct {
	DisplayName *string `json:"display_name,omitempty" binding:"omitempty"`
	Username    string  `json:"username" binding:"required,gte=4"`
	Email       string  `json:"email,omitempty" binding:"omitempty,email"`
	Password    string  `json:"password" binding:"required,gte=6"`
	Role        *string `json:"role,omitempty" binding:"omitempty"`
	TeamID      *int    `json:"team_id,omitempty" binding:"omitempty,numeric"`
}

type RegisterResponse struct {
	User *models.User
}

type LoginRequest struct {
	Username string `json:"username" binding:"omitempty,gte=4"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,gte=6"`
}

type LoginResponse struct {
	User *models.TokenizedUser
}
