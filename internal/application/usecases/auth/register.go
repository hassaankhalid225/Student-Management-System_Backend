package auth

import (
	"context"
	"errors"

	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase interface {
	Execute(ctx context.Context, req dto.RegisterRequest) error
}

type registerUseCaseImpl struct {
	userRepo repositories.UserRepository
}

func NewRegisterUseCase(userRepo repositories.UserRepository) RegisterUseCase {
	return &registerUseCaseImpl{userRepo: userRepo}
}

func (u *registerUseCaseImpl) Execute(ctx context.Context, req dto.RegisterRequest) error {
	existingUser, _ := u.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         req.Role,
		IsActive:     true,
	}

	return u.userRepo.Create(ctx, user)
}
