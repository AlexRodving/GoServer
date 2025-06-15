package task

import "time"

type CreateTaskInput struct {
	Title     string    `json:"title" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Status    string    `json:"status" binding:"omitempty,oneof=todo in-progress done"`
}

type UpdateTaskInput struct {
	Title     *string    `json:"title"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	Status    *string    `json:"status" binding:"omitempty,oneof=todo in-progress done"`
}
