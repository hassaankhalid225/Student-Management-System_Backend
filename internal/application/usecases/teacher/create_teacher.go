package teacher

import (
	"context"

	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateTeacherUseCase interface {
	Execute(ctx context.Context, req dto.CreateTeacherDTO) error
}

type createTeacherUseCaseImpl struct {
	db          *gorm.DB
	teacherRepo repositories.TeacherRepository
}

func NewCreateTeacherUseCase(db *gorm.DB, teacherRepo repositories.TeacherRepository) CreateTeacherUseCase {
	return &createTeacherUseCaseImpl{
		db:          db,
		teacherRepo: teacherRepo,
	}
}

func (u *createTeacherUseCaseImpl) Execute(ctx context.Context, req dto.CreateTeacherDTO) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := &entities.User{
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
			Role:         entities.RoleTeacher,
			IsActive:     true,
		}

		if err := tx.Create(user).Error; err != nil {
			return err
		}

		teacher := &entities.Teacher{
			UserID:         user.ID,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			DateOfBirth:    req.DateOfBirth,
			Gender:         req.Gender,
			Phone:          req.Phone,
			Address:        req.Address,
			Qualification:  req.Qualification,
			Specialization: req.Specialization,
			JoiningDate:    req.JoiningDate,
			Salary:         req.Salary,
		}

		if err := tx.Create(teacher).Error; err != nil {
			return err
		}

		return nil
	})
}
