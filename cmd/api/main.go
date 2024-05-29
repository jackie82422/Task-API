package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"task-api/internal/domain/task"
	"task-api/internal/handlers"
	"task-api/internal/infrastructure/persistence/memory"
	"task-api/internal/infrastructure/persistence/mysql"
	taskService "task-api/internal/services/task"
)

func main() {
	container := buildContainer()

	err := container.Invoke(func(handler *handlers.TaskHandler) {
		r := gin.Default()
		r.GET("/tasks", handler.GetTasks)
		r.GET("/tasks/:id", handler.GetTask)
		r.POST("/tasks", handler.CreateTask)
		r.PUT("/tasks/:id", handler.UpdateTask)
		r.DELETE("/tasks/:id", handler.DeleteTask)

		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to invoke TaskHandler: %v", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "mysql" {
		dsn := os.Getenv("MYSQL_DSN")
		if dsn == "" {
			dsn = "root:root@tcp(localhost:3306)/testDB"
		}
		container.Provide(func() (task.Repository, error) {
			return mysql.NewMySQLTaskRepository(dsn)
		})
	} else {
		container.Provide(func() task.Repository {
			return memory.NewInMemoryTaskRepository()
		})
	}

	container.Provide(taskService.NewTaskService)
	container.Provide(handlers.NewTaskHandler)

	return container
}
