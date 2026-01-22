package student

import (
	"context"

	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/domain/repositories"
)

type UpdateStudentUseCase interface {
	Execute(ctx context.Context, id uuid.UUID, req dto.UpdateStudentDTO) error
}

type updateStudentUseCaseImpl struct {
	studentRepo repositories.StudentRepository
}

func NewUpdateStudentUseCase(studentRepo repositories.StudentRepository) UpdateStudentUseCase {
	return &updateStudentUseCaseImpl{studentRepo: studentRepo}
}

func (u *updateStudentUseCaseImpl) Execute(ctx context.Context, id uuid.UUID, req dto.UpdateStudentDTO) error {
	student, err := u.studentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if req.FirstName != "" {
		student.FirstName = req.FirstName
	}
	if req.LastName != "" {
		student.LastName = req.LastName
	}
	if !req.DateOfBirth.IsZero() {
		student.DateOfBirth = req.DateOfBirth
	}
	if req.Gender != "" {
		student.Gender = req.Gender
	}
	if req.Phone != "" {
		student.Phone = req.Phone
	}
	if req.Address != "" {
		student.Address = req.Address
	}
	if req.ClassID != nil {
		student.ClassID = req.ClassID
	}
	if req.SectionID != nil {
		student.SectionID = req.SectionID
	}
	if req.GuardianName != "" {
		student.GuardianName = req.GuardianName
	}
	if req.GuardianPhone != "" {
		student.GuardianPhone = req.GuardianPhone
	}
	if req.GuardianEmail != "" {
		student.GuardianEmail = req.GuardianEmail
	}
	if req.ProfileImageURL != "" {
		student.ProfileImageURL = req.ProfileImageURL
	}

	return u.studentRepo.Update(ctx, student)
}
