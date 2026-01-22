package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	GradeLevel  int            `json:"grade_level"`
	Sections    []Section      `gorm:"foreignKey:ClassID" json:"sections"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Section struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ClassID   uuid.UUID      `gorm:"type:uuid;not null" json:"class_id"`
	Name      string         `gorm:"not null" json:"name"`
	Capacity  int            `json:"capacity"`
	TeacherID *uuid.UUID     `gorm:"type:uuid" json:"teacher_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
