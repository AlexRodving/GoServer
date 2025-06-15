package database

import (
	"log"

	"server/internal/task"
	"server/internal/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error

	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("База данных подключена.")

	// Миграция моделей
	err = DB.AutoMigrate(&user.User{}, &task.Task{})
	if err != nil {
		return nil, err
	}

	log.Println("Миграция прошла успешно.")
	return DB, nil
}
