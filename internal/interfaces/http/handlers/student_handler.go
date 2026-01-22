package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/application/usecases/student"
)

type StudentHandler struct {
	createStudentUseCase student.CreateStudentUseCase
	getStudentsUseCase    student.GetStudentsUseCase
	getStudentByIDUseCase student.GetStudentByIDUseCase
	updateStudentUseCase student.UpdateStudentUseCase
	deleteStudentUseCase student.DeleteStudentUseCase
}

func NewStudentHandler(
	createStudentUseCase student.CreateStudentUseCase,
	getStudentsUseCase student.GetStudentsUseCase,
	getStudentByIDUseCase student.GetStudentByIDUseCase,
	updateStudentUseCase student.UpdateStudentUseCase,
	deleteStudentUseCase student.DeleteStudentUseCase,
) *StudentHandler {
	return &StudentHandler{
		createStudentUseCase: createStudentUseCase,
		getStudentsUseCase:    getStudentsUseCase,
		getStudentByIDUseCase: getStudentByIDUseCase,
		updateStudentUseCase: updateStudentUseCase,
		deleteStudentUseCase: deleteStudentUseCase,
	}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req dto.CreateStudentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.createStudentUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "student created successfully"})
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	filter := make(map[string]interface{})
	if classID := c.Query("class_id"); classID != "" {
		if uid, err := uuid.Parse(classID); err == nil {
			filter["class_id"] = uid
		}
	}

	students, err := h.getStudentsUseCase.Execute(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	student, err := h.getStudentByIDUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	var req dto.UpdateStudentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.updateStudentUseCase.Execute(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student updated successfully"})
}

func (h *StudentHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	err = h.deleteStudentUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
