package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

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

	err := container.Provide(memory.NewInMemoryTaskRepository)
	if err != nil {
		return nil
	}

	err = container.Provide(handlers.NewTaskHandler)
	if err != nil {
		return nil
	}

	return container
}
