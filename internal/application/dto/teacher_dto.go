package dto

import (
	"time"
)

type CreateTeacherDTO struct {
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" binding:"required,min=6"`
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Qualification  string    `json:"qualification"`
	Specialization string    `json:"specialization"`
	JoiningDate    time.Time `json:"joining_date"`
	Salary         float64   `json:"salary"`
}

type UpdateTeacherDTO struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Qualification  string    `json:"qualification"`
	Specialization string    `json:"specialization"`
	Salary         float64   `json:"salary"`
}
