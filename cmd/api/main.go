package main

import (
	"log"
	"task-api/internal/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"

	"task-api/internal/handlers"
	taskService "task-api/internal/services/task"
)

func main() {
	r := gin.Default()
	injector := config.NewInjector()
	if err := injector.InvokeStorage(); err != nil {
		log.Fatal(err)
	}

	// Define all handler constructors
	handlerConstructors := []interface{}{
		handlers.NewTaskHandler,
	}

	// Register all handlers
	if err := injector.RegisterHandlers(r, handlerConstructors); err != nil {
		log.Fatalf("Failed to register handlers: %v", err)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(taskService.NewTaskService)
	container.Provide(handlers.NewTaskHandler)

	return container
}
