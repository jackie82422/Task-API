package task

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"task-api/internal/domain/task"
	"task-api/internal/mocks"
)

func Test_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewTaskService(mockRepo)

	tests := []struct {
		TestCase string
		Input    task.Info
		Expected task.Info
		Error    error
		Setup    func()
	}{
		{
			TestCase: "Create valid task",
			Input:    task.Info{Name: "Valid Task", Status: 0},
			Expected: task.Info{ID: 1, Name: "Valid Task", Status: 0},
			Error:    nil,
			Setup: func() {
				mockRepo.EXPECT().Create(task.Info{Name: "Valid Task", Status: 0}).Return(task.Info{ID: 1, Name: "Valid Task", Status: 0}, nil)
			},
		},
		{
			TestCase: "Create task with empty name",
			Input:    task.Info{Name: "", Status: 0},
			Expected: task.Info{},
			Error:    errors.New("task name is required"),
			Setup:    func() {},
		},
		{
			TestCase: "Create task with invalid status",
			Input:    task.Info{Name: "Invalid Status Task", Status: 2},
			Expected: task.Info{},
			Error:    errors.New("invalid task status"),
			Setup:    func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			result, err := service.CreateTask(tc.Input)
			if tc.Error != nil {
				assert.Equal(t, tc.Error, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func Test_GetAllTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewTaskService(mockRepo)

	tests := []struct {
		TestCase string
		Expected []task.Info
		Error    error
		Setup    func()
	}{
		{
			TestCase: "Get all tasks",
			Expected: []task.Info{
				{ID: 1, Name: "Test Task 1", Status: 0},
				{ID: 2, Name: "Test Task 2", Status: 1},
			},
			Error: nil,
			Setup: func() {
				mockRepo.EXPECT().GetAll().Return([]task.Info{
					{ID: 1, Name: "Test Task 1", Status: 0},
					{ID: 2, Name: "Test Task 2", Status: 1},
				}, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			result, err := service.GetAllTasks()
			if tc.Error != nil {
				assert.Equal(t, tc.Error, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func Test_GetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewTaskService(mockRepo)

	tests := []struct {
		TestCase string
		ID       int
		Expected task.Info
		Error    error
		Setup    func()
	}{
		{
			TestCase: "Get task by ID",
			ID:       1,
			Expected: task.Info{ID: 1, Name: "Test Task", Status: 0},
			Error:    nil,
			Setup: func() {
				mockRepo.EXPECT().GetByID(1).Return(task.Info{ID: 1, Name: "Test Task", Status: 0}, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			result, err := service.GetTaskByID(tc.ID)
			if tc.Error != nil {
				assert.Equal(t, tc.Error, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func Test_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewTaskService(mockRepo)

	tests := []struct {
		TestCase string
		Input    task.Info
		Expected task.Info
		Error    error
		Setup    func()
	}{
		{
			TestCase: "Update valid task",
			Input:    task.Info{ID: 1, Name: "Updated Task", Status: 1},
			Expected: task.Info{ID: 1, Name: "Updated Task", Status: 1},
			Error:    nil,
			Setup: func() {
				mockRepo.EXPECT().Update(task.Info{ID: 1, Name: "Updated Task", Status: 1}).Return(task.Info{ID: 1, Name: "Updated Task", Status: 1}, nil)
			},
		},
		{
			TestCase: "Update task with empty name",
			Input:    task.Info{ID: 1, Name: "", Status: 1},
			Expected: task.Info{},
			Error:    errors.New("task name is required"),
			Setup:    func() {},
		},
		{
			TestCase: "Update task with invalid status",
			Input:    task.Info{ID: 1, Name: "Invalid Status Task", Status: 2},
			Expected: task.Info{},
			Error:    errors.New("invalid task status"),
			Setup:    func() {},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			result, err := service.UpdateTask(tc.Input)
			if tc.Error != nil {
				assert.Equal(t, tc.Error, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func Test_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepository(ctrl)
	service := NewTaskService(mockRepo)

	tests := []struct {
		TestCase string
		ID       int
		Error    error
		Setup    func()
	}{
		{
			TestCase: "Delete task",
			ID:       1,
			Error:    nil,
			Setup: func() {
				mockRepo.EXPECT().Delete(1).Return(nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			err := service.DeleteTask(tc.ID)
			if tc.Error != nil {
				assert.Equal(t, tc.Error, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
