package auth

import (
	"context"
	"errors"

	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"github.com/yourusername/sms-backend/internal/infrastructure/auth"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Execute(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type loginUseCaseImpl struct {
	userRepo   repositories.UserRepository
	jwtService *auth.JWTService
}

func NewLoginUseCase(userRepo repositories.UserRepository, jwtService *auth.JWTService) LoginUseCase {
	return &loginUseCaseImpl{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (u *loginUseCaseImpl) Execute(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := u.jwtService.GenerateTokens(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}
