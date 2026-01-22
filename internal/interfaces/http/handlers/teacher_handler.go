package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/application/usecases/teacher"
)

type TeacherHandler struct {
	createTeacherUseCase teacher.CreateTeacherUseCase
	getTeachersUseCase    teacher.GetTeachersUseCase
}

func NewTeacherHandler(
	createTeacherUseCase teacher.CreateTeacherUseCase,
	getTeachersUseCase teacher.GetTeachersUseCase,
) *TeacherHandler {
	return &TeacherHandler{
		createTeacherUseCase: createTeacherUseCase,
		getTeachersUseCase:    getTeachersUseCase,
	}
}

func (h *TeacherHandler) Create(c *gin.Context) {
	var req dto.CreateTeacherDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.createTeacherUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "teacher created successfully"})
}

func (h *TeacherHandler) GetAll(c *gin.Context) {
	teachers, err := h.getTeachersUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teachers)
}
