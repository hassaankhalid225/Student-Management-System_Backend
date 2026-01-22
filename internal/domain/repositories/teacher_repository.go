package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/entities"
)

type TeacherRepository interface {
	Create(ctx context.Context, teacher *entities.Teacher) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Teacher, error)
	FindAll(ctx context.Context) ([]*entities.Teacher, error)
	Update(ctx context.Context, teacher *entities.Teacher) error
	Delete(ctx context.Context, id uuid.UUID) error
}
