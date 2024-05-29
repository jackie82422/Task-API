package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"task-api/internal/domain/task"
	"task-api/internal/mocks"
)

func TestTaskHandler_GetTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks", handler.GetTasks)

	tests := []struct {
		TestCase string
		Expected []task.Info
		Setup    func()
		Status   int
	}{
		{
			TestCase: "Get all tasks",
			Expected: []task.Info{
				{ID: 1, Name: "Test Task 1", Status: 0},
				{ID: 2, Name: "Test Task 2", Status: 1},
			},
			Setup: func() {
				mockService.EXPECT().GetAllTasks().Return([]task.Info{
					{ID: 1, Name: "Test Task 1", Status: 0},
					{ID: 2, Name: "Test Task 2", Status: 1},
				}, nil)
			},
			Status: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.Status, rr.Code)
			var response []task.Info
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, response)
		})
	}
}

func TestTaskHandler_GetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/tasks/:id", handler.GetTask)

	tests := []struct {
		TestCase string
		ID       int
		Expected task.Info
		Setup    func()
		Status   int
	}{
		{
			TestCase: "Get task by ID",
			ID:       1,
			Expected: task.Info{ID: 1, Name: "Test Task", Status: 0},
			Setup: func() {
				mockService.EXPECT().GetTaskByID(1).Return(task.Info{ID: 1, Name: "Test Task", Status: 0}, nil)
			},
			Status: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			req, _ := http.NewRequest(http.MethodGet, "/tasks/"+strconv.Itoa(tc.ID), nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.Status, rr.Code)
			var response task.Info
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, response)
		})
	}
}

func TestTaskHandler_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/tasks", handler.CreateTask)

	tests := []struct {
		TestCase string
		Input    task.Info
		Expected task.Info
		Setup    func()
		Status   int
	}{
		{
			TestCase: "Create valid task",
			Input:    task.Info{Name: "New Task", Status: 0},
			Expected: task.Info{ID: 1, Name: "New Task", Status: 0},
			Setup: func() {
				mockService.EXPECT().CreateTask(task.Info{Name: "New Task", Status: 0}).Return(task.Info{ID: 1, Name: "New Task", Status: 0}, nil)
			},
			Status: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			jsonValue, _ := json.Marshal(tc.Input)
			req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.Status, rr.Code)
			var response task.Info
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, response)
		})
	}
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/tasks/:id", handler.UpdateTask)

	tests := []struct {
		TestCase string
		ID       int
		Input    task.Info
		Expected task.Info
		Setup    func()
		Status   int
	}{
		{
			TestCase: "Update valid task",
			ID:       1,
			Input:    task.Info{Name: "Updated Task", Status: 1},
			Expected: task.Info{ID: 1, Name: "Updated Task", Status: 1},
			Setup: func() {
				mockService.EXPECT().UpdateTask(task.Info{ID: 1, Name: "Updated Task", Status: 1}).Return(task.Info{ID: 1, Name: "Updated Task", Status: 1}, nil)
			},
			Status: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			jsonValue, _ := json.Marshal(tc.Input)
			req, _ := http.NewRequest(http.MethodPut, "/tasks/"+strconv.Itoa(tc.ID), bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.Status, rr.Code)
			var response task.Info
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, response)
		})
	}
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := NewTaskHandler(mockService)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/tasks/:id", handler.DeleteTask)

	tests := []struct {
		TestCase string
		ID       int
		Setup    func()
		Status   int
	}{
		{
			TestCase: "Delete task",
			ID:       1,
			Setup: func() {
				mockService.EXPECT().DeleteTask(1).Return(nil)
			},
			Status: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tc.Setup()
			req, _ := http.NewRequest(http.MethodDelete, "/tasks/"+strconv.Itoa(tc.ID), nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, tc.Status, rr.Code)
			var response map[string]string
			err := json.Unmarshal(rr.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, "Task deleted successfully", response["message"])
		})
	}
}
