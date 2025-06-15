package main

import (
	//"fmt"
	"log"
	//"os"
	"server/config"
	"github.com/gin-gonic/gin"
	"server/internal/database"
	"server/internal/auth"
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

	// Запускаем сервер
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
