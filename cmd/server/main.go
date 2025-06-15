package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/config"
	"server/internal/auth"
	"server/internal/database"
	"server/internal/task"
)

func main() {
	// Загрузка .env
	config.LoadConfig()

	// Получаем порт из окружения
	port := config.GetEnv("PORT", "8080")

	// Инициализируем БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Ошибка при инициализации БД: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Создаем Gin-роутер
	r := gin.Default()

	authHandler := auth.NewAuthHandler(db)

	// Базовый маршрут
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Группа auth
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}
	taskHandler := task.NewTaskHandler(db)
	api := r.Group("/api")
	api.Use(auth.AuthMiddleware())
	{
		api.POST("/tasks", taskHandler.CreateTask)
		api.GET("/tasks", taskHandler.GetTasks)
		api.PUT("/tasks/:id", taskHandler.UpdateTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}
	// Запускаем сервер
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
