package task

import (
	"errors"

	"task-api/internal/domain/task"
)

type taskService struct {
	repo task.Repository
}

func NewTaskService(repo task.Repository) Service {
	return &taskService{repo: repo}
}

func (s *taskService) GetAllTasks() ([]task.Info, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTaskByID(id int) (task.Info, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) CreateTask(t task.Info) (task.Info, error) {
	if t.Name == "" {
		return task.Info{}, errors.New("task name is required")
	}
	if t.Status != 0 && t.Status != 1 {
		return task.Info{}, errors.New("invalid task status")
	}
	return s.repo.Create(t)
}

func (s *taskService) UpdateTask(t task.Info) (task.Info, error) {
	if t.Name == "" {
		return task.Info{}, errors.New("task name is required")
	}
	if t.Status != 0 && t.Status != 1 {
		return task.Info{}, errors.New("invalid task status")
	}
	return s.repo.Update(t)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
