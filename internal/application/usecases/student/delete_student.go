package student

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
)

type DeleteStudentUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) error
}

type deleteStudentUseCaseImpl struct {
	studentRepo repositories.StudentRepository
}

func NewDeleteStudentUseCase(studentRepo repositories.StudentRepository) DeleteStudentUseCase {
	return &deleteStudentUseCaseImpl{studentRepo: studentRepo}
}

func (u *deleteStudentUseCaseImpl) Execute(ctx context.Context, id uuid.UUID) error {
	return u.studentRepo.Delete(ctx, id)
}
