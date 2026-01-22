package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type teacherRepositoryImpl struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) repositories.TeacherRepository {
	return &teacherRepositoryImpl{db: db}
}

func (r *teacherRepositoryImpl) Create(ctx context.Context, teacher *entities.Teacher) error {
	return r.db.WithContext(ctx).Create(teacher).Error
}

func (r *teacherRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entities.Teacher, error) {
	var teacher entities.Teacher
	if err := r.db.WithContext(ctx).Preload("User").First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *teacherRepositoryImpl) FindAll(ctx context.Context) ([]*entities.Teacher, error) {
	var teachers []*entities.Teacher
	if err := r.db.WithContext(ctx).Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *teacherRepositoryImpl) Update(ctx context.Context, teacher *entities.Teacher) error {
	return r.db.WithContext(ctx).Save(teacher).Error
}

func (r *teacherRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Teacher{}, id).Error
}
