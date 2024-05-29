package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"os"
	"task-api/internal/domain/task"
	"task-api/internal/infrastructure/persistence/mysql"

	"task-api/internal/handlers"
	"task-api/internal/infrastructure/persistence/memory"
)

func main() {
	container := buildContainer()
	err := container.Invoke(func(handler *handlers.TaskHandler) {
		r := gin.Default()

		r.GET("/tasks", handler.GetTasks)
		r.POST("/tasks", handler.CreateTask)
		r.PUT("/tasks/:id", handler.UpdateTask)
		r.DELETE("/tasks/:id", handler.DeleteTask)

		err := r.Run()
		if err != nil {
			return
		}
	})

	if err != nil {
		panic(err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "mysql" {
		dsn := os.Getenv("MYSQL_DSN")
		container.Provide(func() (task.Repository, error) {
			return mysql.NewMySQLTaskRepository(dsn)
		})
	} else {
		container.Provide(memory.NewInMemoryTaskRepository)
	}

	container.Provide(handlers.NewTaskHandler)

	return container
}
