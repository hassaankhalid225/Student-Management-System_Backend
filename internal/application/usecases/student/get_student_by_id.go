package student

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
)

type GetStudentByIDUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*entities.Student, error)
}

type getStudentByIDUseCaseImpl struct {
	studentRepo repositories.StudentRepository
}

func NewGetStudentByIDUseCase(studentRepo repositories.StudentRepository) GetStudentByIDUseCase {
	return &getStudentByIDUseCaseImpl{studentRepo: studentRepo}
}

func (u *getStudentByIDUseCaseImpl) Execute(ctx context.Context, id uuid.UUID) (*entities.Student, error) {
	return u.studentRepo.FindByID(ctx, id)
}
