package task

import "gorm.io/gorm"

type Repository interface {
	Create(task *Task) error
	FindAllByUser(userID uint) ([]Task, error)
	FindByID(id, userID uint) (*Task, error)
	Update(task *Task) error
	Delete(id, userID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(task *Task) error {
	return r.db.Create(task).Error
}

func (r *repository) FindAllByUser(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *repository) FindByID(id, userID uint) (*Task, error) {
	var task Task
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *repository) Update(task *Task) error {
	return r.db.Save(task).Error
}

func (r *repository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&Task{}).Error
}
