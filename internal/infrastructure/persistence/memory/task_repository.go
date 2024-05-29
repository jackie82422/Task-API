package memory

import (
	"errors"
	"sync"

	"task-api/internal/domain/task"
)

type TaskRepository struct {
	mu     sync.Mutex
	tasks  map[int]task.Task
	nextID int
}

func NewInMemoryTaskRepository() task.Repository {
	return &TaskRepository{
		tasks:  make(map[int]task.Task),
		nextID: 1,
	}
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []task.Task
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *TaskRepository) GetByID(id int) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, exists := r.tasks[id]
	if !exists {
		return task.Task{}, errors.New("task not found")
	}
	return t, nil
}

func (r *TaskRepository) Create(t task.Task) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks[t.ID] = t
	return t, nil
}

func (r *TaskRepository) Update(t task.Task) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.tasks[t.ID]
	if !exists {
		return task.Task{}, errors.New("task not found")
	}
	r.tasks[t.ID] = t
	return t, nil
}

func (r *TaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
