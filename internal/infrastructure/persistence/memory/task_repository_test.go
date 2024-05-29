package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"task-api/internal/domain/task"
)

func TestCreateTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	tests := []struct {
		TestCase string
		Input    task.Task
		Expected task.Task
	}{
		{
			TestCase: "Create Task",
			Input:    task.Task{Name: "Test Task", Status: 0},
			Expected: task.Task{ID: 1, Name: "Test Task", Status: 0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			createdTask, err := repo.Create(tc.Input)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected.ID, createdTask.ID)
			assert.Equal(t, tc.Expected.Name, createdTask.Name)
			assert.Equal(t, tc.Expected.Status, createdTask.Status)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	newTask := task.Task{Name: "Test Task", Status: 0}
	repo.Create(newTask)

	tests := []struct {
		TestCase string
		Expected []task.Task
	}{
		{
			TestCase: "Get All Tasks",
			Expected: []task.Task{
				{ID: 1, Name: "Test Task", Status: 0},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tasks, err := repo.GetAll()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, tasks)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	newTask := task.Task{Name: "Test Task", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		ID       int
		Expected task.Task
	}{
		{
			TestCase: "Get Task By ID",
			ID:       createdTask.ID,
			Expected: createdTask,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			task, err := repo.GetByID(tc.ID)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, task)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	newTask := task.Task{Name: "Test Task", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		Input    task.Task
		Expected task.Task
	}{
		{
			TestCase: "Update Task",
			Input:    task.Task{ID: createdTask.ID, Name: "Updated Task", Status: 1},
			Expected: task.Task{ID: createdTask.ID, Name: "Updated Task", Status: 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			updatedTask, err := repo.Update(tc.Input)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected.ID, updatedTask.ID)
			assert.Equal(t, tc.Expected.Name, updatedTask.Name)
			assert.Equal(t, tc.Expected.Status, updatedTask.Status)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	repo := NewInMemoryTaskRepository()

	newTask := task.Task{Name: "Test Task", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		ID       int
	}{
		{
			TestCase: "Delete Task",
			ID:       createdTask.ID,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			err := repo.Delete(tc.ID)
			assert.NoError(t, err)

			tasks, err := repo.GetAll()
			assert.NoError(t, err)
			assert.Empty(t, tasks)
		})
	}
}
