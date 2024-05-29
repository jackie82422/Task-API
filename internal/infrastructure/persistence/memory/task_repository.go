package memory

import (
	"errors"
	"sync"

	"task-api/internal/domain/task"
)

type TaskRepository struct {
	mu     sync.Mutex
	tasks  map[int]task.Info
	nextID int
}

func NewInMemoryTaskRepository() task.Repository {
	return &TaskRepository{
		tasks:  make(map[int]task.Info),
		nextID: 1,
	}
}

func (r *TaskRepository) GetAll() ([]task.Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []task.Info
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *TaskRepository) GetByID(id int) (task.Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, exists := r.tasks[id]
	if !exists {
		return task.Info{}, errors.New("task not found")
	}
	return t, nil
}

func (r *TaskRepository) Create(t task.Info) (task.Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t.ID = r.nextID
	r.nextID++
	r.tasks[t.ID] = t
	return t, nil
}

func (r *TaskRepository) Update(t task.Info) (task.Info, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.tasks[t.ID]
	if !exists {
		return task.Info{}, errors.New("task not found")
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
