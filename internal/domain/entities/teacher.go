package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	FirstName       string         `gorm:"not null" json:"first_name"`
	LastName        string         `gorm:"not null" json:"last_name"`
	DateOfBirth     time.Time      `json:"date_of_birth"`
	Gender          string         `json:"gender"`
	Phone           string         `json:"phone"`
	Address         string         `json:"address"`
	Qualification   string         `json:"qualification"`
	Specialization  string         `json:"specialization"`
	JoiningDate     time.Time      `json:"joining_date"`
	Salary          float64        `json:"salary"`
	ProfileImageURL string         `json:"profile_image_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type Subject struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Code        string         `gorm:"uniqueIndex" json:"code"`
	Description string         `json:"description"`
	ClassID     uuid.UUID      `gorm:"type:uuid;not null" json:"class_id"`
	Class       Class          `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	TeacherID   *uuid.UUID     `gorm:"type:uuid" json:"teacher_id"`
	Teacher     *Teacher       `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
