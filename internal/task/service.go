package task

type Service interface {
	CreateTask(input CreateTaskInput, userID uint) error
	GetTasks(userID uint) ([]Task, error)
	UpdateTask(id uint, input UpdateTaskInput, userID uint) error
	DeleteTask(id uint, userID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateTask(input CreateTaskInput, userID uint) error {
	task := Task{
		Title:     input.Title,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Status:    input.Status,
		UserID:    userID,
	}
	return s.repo.Create(&task)
}

func (s *service) GetTasks(userID uint) ([]Task, error) {
	return s.repo.FindAllByUser(userID)
}

func (s *service) UpdateTask(id uint, input UpdateTaskInput, userID uint) error {
	task, err := s.repo.FindByID(id, userID)
	if err != nil {
		return err
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.StartDate != nil {
		task.StartDate = *input.StartDate
	}
	if input.EndDate != nil {
		task.EndDate = *input.EndDate
	}
	if input.Status != nil {
		task.Status = *input.Status
	}

	return s.repo.Update(task)
}

func (s *service) DeleteTask(id uint, userID uint) error {
	return s.repo.Delete(id, userID)
}
