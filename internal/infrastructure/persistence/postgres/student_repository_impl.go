package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type studentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) repositories.StudentRepository {
	return &studentRepositoryImpl{db: db}
}

func (r *studentRepositoryImpl) Create(ctx context.Context, student *entities.Student) error {
	return r.db.WithContext(ctx).Create(student).Error
}

func (r *studentRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entities.Student, error) {
	var student entities.Student
	if err := r.db.WithContext(ctx).Preload("User").Preload("Class").Preload("Section").First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepositoryImpl) FindAll(ctx context.Context, filter map[string]interface{}) ([]*entities.Student, error) {
	var students []*entities.Student
	query := r.db.WithContext(ctx).Preload("Class").Preload("Section")
	
	if classID, ok := filter["class_id"]; ok {
		query = query.Where("class_id = ?", classID)
	}
	
	if err := query.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepositoryImpl) Update(ctx context.Context, student *entities.Student) error {
	return r.db.WithContext(ctx).Save(student).Error
}

func (r *studentRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entities.Student{}, id).Error
}
