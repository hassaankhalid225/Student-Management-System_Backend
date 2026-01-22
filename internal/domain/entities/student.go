package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	FirstName       string         `gorm:"not null" json:"first_name"`
	LastName        string         `gorm:"not null" json:"last_name"`
	DateOfBirth     time.Time      `json:"date_of_birth"`
	Gender          string         `json:"gender"`
	Phone           string         `json:"phone"`
	Address         string         `json:"address"`
	ClassID         *uuid.UUID     `gorm:"type:uuid" json:"class_id"`
	Class           *Class         `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	SectionID       *uuid.UUID     `gorm:"type:uuid" json:"section_id"`
	Section         *Section       `gorm:"foreignKey:SectionID" json:"section,omitempty"`
	RollNumber      string         `gorm:"uniqueIndex" json:"roll_number"`
	AdmissionDate   time.Time      `json:"admission_date"`
	GuardianName    string         `json:"guardian_name"`
	GuardianPhone   string         `json:"guardian_phone"`
	GuardianEmail   string         `json:"guardian_email"`
	ProfileImageURL string         `json:"profile_image_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}