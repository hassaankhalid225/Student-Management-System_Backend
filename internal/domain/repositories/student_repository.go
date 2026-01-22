package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/entities"
)

type StudentRepository interface {
	Create(ctx context.Context, student *entities.Student) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Student, error)
	FindAll(ctx context.Context, filter map[string]interface{}) ([]*entities.Student, error)
	Update(ctx context.Context, student *entities.Student) error
	Delete(ctx context.Context, id uuid.UUID) error
}
