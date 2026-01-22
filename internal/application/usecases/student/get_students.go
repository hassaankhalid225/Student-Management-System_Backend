package student

import (
	"context"

	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
)

type GetStudentsUseCase interface {
	Execute(ctx context.Context, filter map[string]interface{}) ([]*entities.Student, error)
}

type getStudentsUseCaseImpl struct {
	studentRepo repositories.StudentRepository
}

func NewGetStudentsUseCase(studentRepo repositories.StudentRepository) GetStudentsUseCase {
	return &getStudentsUseCaseImpl{studentRepo: studentRepo}
}

func (u *getStudentsUseCaseImpl) Execute(ctx context.Context, filter map[string]interface{}) ([]*entities.Student, error) {
	return u.studentRepo.FindAll(ctx, filter)
}
