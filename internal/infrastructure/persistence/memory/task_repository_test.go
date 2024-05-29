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
		Input    task.Info
		Expected task.Info
	}{
		{
			TestCase: "Create TaskInfo",
			Input:    task.Info{Name: "Test TaskInfo", Status: 0},
			Expected: task.Info{ID: 1, Name: "Test TaskInfo", Status: 0},
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

	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	repo.Create(newTask)

	tests := []struct {
		TestCase string
		Expected []task.Info
	}{
		{
			TestCase: "Get All Tasks",
			Expected: []task.Info{
				{ID: 1, Name: "Test TaskInfo", Status: 0},
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

	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		ID       int
		Expected task.Info
	}{
		{
			TestCase: "Get TaskInfo By ID",
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

	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		Input    task.Info
		Expected task.Info
	}{
		{
			TestCase: "Update TaskInfo",
			Input:    task.Info{ID: createdTask.ID, Name: "Updated TaskInfo", Status: 1},
			Expected: task.Info{ID: createdTask.ID, Name: "Updated TaskInfo", Status: 1},
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

	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	createdTask, _ := repo.Create(newTask)

	tests := []struct {
		TestCase string
		ID       int
	}{
		{
			TestCase: "Delete TaskInfo",
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
