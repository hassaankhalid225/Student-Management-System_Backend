package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/sms-backend/internal/config"
	"github.com/yourusername/sms-backend/internal/domain/entities"
	"github.com/yourusername/sms-backend/internal/infrastructure/auth"
	"github.com/yourusername/sms-backend/internal/infrastructure/persistence/postgres"
	auth_usecases "github.com/yourusername/sms-backend/internal/application/usecases/auth"
	student_usecases "github.com/yourusername/sms-backend/internal/application/usecases/student"
	teacher_usecases "github.com/yourusername/sms-backend/internal/application/usecases/teacher"
	"github.com/yourusername/sms-backend/internal/interfaces/http/handlers"
	"github.com/yourusername/sms-backend/internal/interfaces/http/middleware"
)

func main() {
	// ... (config and db init)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db := config.InitDB(cfg)

	// Run migrations
	err = db.AutoMigrate(&entities.User{}, &entities.Class{}, &entities.Section{}, &entities.Student{}, &entities.Teacher{}, &entities.Subject{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Services
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.RefreshSecret)

	// Initialize Repositories
	userRepo := postgres.NewUserRepository(db)
	studentRepo := postgres.NewStudentRepository(db)
	teacherRepo := postgres.NewTeacherRepository(db)

	// Initialize Use Cases
	loginUseCase := auth_usecases.NewLoginUseCase(userRepo, jwtService)
	registerUseCase := auth_usecases.NewRegisterUseCase(userRepo)

	createStudentUseCase := student_usecases.NewCreateStudentUseCase(db, studentRepo, userRepo)
	getStudentsUseCase := student_usecases.NewGetStudentsUseCase(studentRepo)
	getStudentByIDUseCase := student_usecases.NewGetStudentByIDUseCase(studentRepo)
	updateStudentUseCase := student_usecases.NewUpdateStudentUseCase(studentRepo)
	deleteStudentUseCase := student_usecases.NewDeleteStudentUseCase(studentRepo)

	createTeacherUseCase := teacher_usecases.NewCreateTeacherUseCase(db, teacherRepo)
	getTeachersUseCase := teacher_usecases.NewGetTeachersUseCase(teacherRepo)

	// Initialize Handlers
	authHandler := handlers.NewAuthHandler(loginUseCase, registerUseCase)
	studentHandler := handlers.NewStudentHandler(
		createStudentUseCase,
		getStudentsUseCase,
		getStudentByIDUseCase,
		updateStudentUseCase,
		deleteStudentUseCase,
	)
	teacherHandler := handlers.NewTeacherHandler(createTeacherUseCase, getTeachersUseCase)

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Router
	r := gin.Default()

	// Routes
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/register", authHandler.Register)
		}

		studentGroup := v1.Group("/students")
		studentGroup.Use(middleware.AuthMiddleware(jwtService))
		{
			studentGroup.POST("", middleware.RoleMiddleware(entities.RoleAdmin), studentHandler.Create)
			studentGroup.GET("", middleware.RoleMiddleware(entities.RoleAdmin, entities.RoleTeacher), studentHandler.GetAll)
			studentGroup.GET("/:id", studentHandler.GetByID)
			studentGroup.PUT("/:id", middleware.RoleMiddleware(entities.RoleAdmin), studentHandler.Update)
			studentGroup.DELETE("/:id", middleware.RoleMiddleware(entities.RoleAdmin), studentHandler.Delete)
		}

		teacherGroup := v1.Group("/teachers")
		teacherGroup.Use(middleware.AuthMiddleware(jwtService))
		{
			teacherGroup.POST("", middleware.RoleMiddleware(entities.RoleAdmin), teacherHandler.Create)
			teacherGroup.GET("", middleware.RoleMiddleware(entities.RoleAdmin), teacherHandler.GetAll)
		}
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// Run server
	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
