package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"task-api/internal/domain/task"
)

const (
	mysqlPassword = "root"
	mysqlDB       = "testDB"
)

var (
	db  *sql.DB
	dsn string
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": mysqlPassword,
			"MYSQL_DATABASE":      mysqlDB,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL").WithStartupTimeout(60 * time.Second),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		fmt.Printf("Failed to start mysql container: %v", err)
		os.Exit(1)
	}

	host, err := mysqlC.Host(ctx)
	if err != nil {
		fmt.Printf("Failed to get mysql container host: %v", err)
		os.Exit(1)
	}

	port, err := mysqlC.MappedPort(ctx, "3306/tcp")
	if err != nil {
		fmt.Printf("Failed to get mysql container port: %v", err)
		os.Exit(1)
	}

	dsn = fmt.Sprintf("root:root@tcp(%s:%s)/testDB", host, port.Port())
	db, err = setupTestDB()
	if err != nil {
		fmt.Printf("Failed to set up test DB: %v", err)
		os.Exit(1)
	}

	code := m.Run()

	err = mysqlC.Terminate(ctx)
	if err != nil {
		fmt.Printf("Failed to terminate test DB: %v", err)
	}

	os.Exit(code)
}

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Create tasks table for testing
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            status INT NOT NULL
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func clearTestDB(db *sql.DB) error {
	_, err := db.Exec(`TRUNCATE TABLE tasks`)
	return err
}

func TestCreateTask(t *testing.T) {
	repo := &TaskRepository{DB: db}

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
			err := clearTestDB(db)
			assert.NoError(t, err)

			createdTask, err := repo.Create(tc.Input)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected.ID, createdTask.ID)
			assert.Equal(t, tc.Expected.Name, createdTask.Name)
			assert.Equal(t, tc.Expected.Status, createdTask.Status)
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	repo := &TaskRepository{DB: db}

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

	err := clearTestDB(db)
	assert.NoError(t, err)

	// First, create a task
	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	_, _ = repo.Create(newTask)

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			tasks, err := repo.GetAll()
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, tasks)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	repo := &TaskRepository{DB: db}

	tests := []struct {
		TestCase string
		ID       int
		Expected task.Info
	}{
		{
			TestCase: "Get TaskInfo By ID",
			ID:       1,
			Expected: task.Info{ID: 1, Name: "Test TaskInfo", Status: 0},
		},
	}

	err := clearTestDB(db)
	assert.NoError(t, err)

	// First, create a task
	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	_, err = repo.Create(newTask)
	assert.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.TestCase, func(t *testing.T) {
			testTask, err := repo.GetByID(tc.ID)
			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, testTask)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	repo := &TaskRepository{DB: db}

	tests := []struct {
		TestCase string
		Input    task.Info
		Expected task.Info
	}{
		{
			TestCase: "Update TaskInfo",
			Input:    task.Info{ID: 1, Name: "Updated TaskInfo", Status: 1},
			Expected: task.Info{ID: 1, Name: "Updated TaskInfo", Status: 1},
		},
	}

	err := clearTestDB(db)
	assert.NoError(t, err)

	// First, create a task
	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	_, err = repo.Create(newTask)
	assert.NoError(t, err)

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
	repo := &TaskRepository{DB: db}

	tests := []struct {
		TestCase string
		ID       int
	}{
		{
			TestCase: "Delete TaskInfo",
			ID:       1,
		},
	}

	err := clearTestDB(db)
	assert.NoError(t, err)

	// First, create a task
	newTask := task.Info{Name: "Test TaskInfo", Status: 0}
	_, err = repo.Create(newTask)
	assert.NoError(t, err)

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
