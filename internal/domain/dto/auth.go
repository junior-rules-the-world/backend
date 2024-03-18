package dto

import "events-organizator/internal/domain/models"

type RegisterRequest struct {
	DisplayName *string `json:"display_name,omitempty" binding:"-"`
	Username    string  `json:"username" binding:"required"`
	Email       string  `json:"email,omitempty" binding:"-"`
	Password    string  `json:"password" binding:"required"`
	Role        *string `json:"role,omitempty" binding:"-"`
	TeamID      *int    `json:"team_id,omitempty" binding:"-"`
}

type RegisterResponse struct {
	User *models.User
}
