package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/sms-backend/internal/application/dto"
	"github.com/yourusername/sms-backend/internal/application/usecases/auth"
)

type AuthHandler struct {
	loginUseCase    auth.LoginUseCase
	registerUseCase auth.RegisterUseCase
}

func NewAuthHandler(loginUseCase auth.LoginUseCase, registerUseCase auth.RegisterUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase:    loginUseCase,
		registerUseCase: registerUseCase,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.loginUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.registerUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}
