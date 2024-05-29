package task

import "task-api/internal/domain/task"

type Service interface {
	GetAllTasks() ([]task.Info, error)
	GetTaskByID(id int) (task.Info, error)
	CreateTask(t task.Info) (task.Info, error)
	UpdateTask(t task.Info) (task.Info, error)
	DeleteTask(id int) error
}
