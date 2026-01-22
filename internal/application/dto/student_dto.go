package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateStudentDTO struct {
	Email         string     `json:"email" binding:"required,email"`
	Password      string     `json:"password" binding:"required,min=6"`
	FirstName     string     `json:"first_name" binding:"required"`
	LastName      string     `json:"last_name" binding:"required"`
	DateOfBirth   time.Time  `json:"date_of_birth"`
	Gender        string     `json:"gender"`
	Phone         string     `json:"phone"`
	Address       string     `json:"address"`
	ClassID       *uuid.UUID `json:"class_id"`
	SectionID     *uuid.UUID `json:"section_id"`
	RollNumber    string     `json:"roll_number" binding:"required"`
	AdmissionDate time.Time  `json:"admission_date"`
	GuardianName  string     `json:"guardian_name"`
	GuardianPhone string     `json:"guardian_phone"`
	GuardianEmail string     `json:"guardian_email"`
}

type UpdateStudentDTO struct {
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	DateOfBirth     time.Time  `json:"date_of_birth"`
	Gender          string     `json:"gender"`
	Phone           string     `json:"phone"`
	Address         string     `json:"address"`
	ClassID         *uuid.UUID `json:"class_id"`
	SectionID       *uuid.UUID `json:"section_id"`
	GuardianName    string     `json:"guardian_name"`
	GuardianPhone   string     `json:"guardian_phone"`
	GuardianEmail   string     `json:"guardian_email"`
	ProfileImageURL string     `json:"profile_image_url"`
}
