package dto

import "github.com/yourusername/sms-backend/internal/domain/entities"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
	User         *entities.User `json:"user"`
}

type RegisterRequest struct {
	Email    string            `json:"email" binding:"required,email"`
	Password string            `json:"password" binding:"required,min=6"`
	Role     entities.UserRole `json:"role" binding:"required"`
}
