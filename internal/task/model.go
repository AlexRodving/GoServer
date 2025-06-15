package task

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Completed bool      `gorm:"default:false" json:"completed"`
	UserID    uint      `gorm:"index" json:"user_id"` // внешний ключ к пользователю
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
