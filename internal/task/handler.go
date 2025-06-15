package task

import (
	"net/http"
	"strconv"

	"server/pkg/context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewTaskHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return &Handler{service}
}

func (h *Handler) CreateTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Некорректные данные",
			"details": err.Error(),
		})
		return
	}

	// Validate dates
	if input.EndDate.Before(input.StartDate) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Дата окончания не может быть раньше даты начала",
		})
		return
	}

	userID := context.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Пользователь не авторизован",
		})
		return
	}

	err := h.service.CreateTask(input, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Ошибка при создании задачи",
			"details": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetTasks(c *gin.Context) {
	userID := context.GetUserID(c)
	tasks, err := h.service.GetTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении задач"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	taskID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	userID := context.GetUserID(c)
	if err := h.service.UpdateTask(uint(taskID), input, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	taskID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	userID := context.GetUserID(c)
	if err := h.service.DeleteTask(uint(taskID), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}
	c.Status(http.StatusOK)
}
