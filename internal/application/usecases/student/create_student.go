package student

import (
	"context"

	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateStudentUseCase interface {
	Execute(ctx context.Context, req dto.CreateStudentDTO) error
}

type createStudentUseCaseImpl struct {
	db          *gorm.DB
	studentRepo repositories.StudentRepository
	userRepo    repositories.UserRepository
}

func NewCreateStudentUseCase(db *gorm.DB, studentRepo repositories.StudentRepository, userRepo repositories.UserRepository) CreateStudentUseCase {
	return &createStudentUseCaseImpl{
		db:          db,
		studentRepo: studentRepo,
		userRepo:    userRepo,
	}
}

func (u *createStudentUseCaseImpl) Execute(ctx context.Context, req dto.CreateStudentDTO) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		// 1. Create User
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := &entities.User{
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
			Role:         entities.RoleStudent,
			IsActive:     true,
		}

		// Use transaction-scoped repository
		// We'll just use GORM directly for simplicity in the transaction or pass TX to repos
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 2. Create Student
		student := &entities.Student{
			UserID:        user.ID,
			FirstName:     req.FirstName,
			LastName:      req.LastName,
			DateOfBirth:   req.DateOfBirth,
			Gender:        req.Gender,
			Phone:         req.Phone,
			Address:       req.Address,
			ClassID:       req.ClassID,
			SectionID:     req.SectionID,
			RollNumber:    req.RollNumber,
			AdmissionDate: req.AdmissionDate,
			GuardianName:  req.GuardianName,
			GuardianPhone: req.GuardianPhone,
			GuardianEmail: req.GuardianEmail,
		}

		if err := tx.Create(student).Error; err != nil {
			return err
		}

		return nil
	})
}
