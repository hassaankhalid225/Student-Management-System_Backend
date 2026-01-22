package teacher

import (
	"context"

	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
)

type GetTeachersUseCase interface {
	Execute(ctx context.Context) ([]*entities.Teacher, error)
}

type getTeachersUseCaseImpl struct {
	teacherRepo repositories.TeacherRepository
}

func NewGetTeachersUseCase(teacherRepo repositories.TeacherRepository) GetTeachersUseCase {
	return &getTeachersUseCaseImpl{teacherRepo: teacherRepo}
}

func (u *getTeachersUseCaseImpl) Execute(ctx context.Context) ([]*entities.Teacher, error) {
	return u.teacherRepo.FindAll(ctx)
}
