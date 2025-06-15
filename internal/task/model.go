package task

import (
	"time"
)

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Status    string    `gorm:"default:'todo'"` // todo, in-progress, done
	UserID    uint      `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
