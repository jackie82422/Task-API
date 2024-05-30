package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"task-api/internal/config"
	"task-api/internal/handlers"
	"task-api/internal/infrastructure/persistence/memory"
	"task-api/internal/infrastructure/persistence/mysql"
	taskService "task-api/internal/services/task"
)

func main() {
	injector := config.NewInjector()

	// Define in-memory and relational repository constructors
	inMemoryRepos := []interface{}{
		memory.NewInMemoryTaskRepository,
	}

	relationRepos := []interface{}{
		mysql.NewMySQLTaskRepository,
	}

	services := []interface{}{
		taskService.NewTaskService,
	}

	handler := []interface{}{
		handlers.NewTaskHandler,
	}

	err := injector.Provide(func() *gin.Engine {
		return gin.Default()
	})
	if err != nil {
		log.Fatal(err)
	}

	// Provide storage injection
	if err := injector.ProvideStorage(inMemoryRepos, relationRepos); err != nil {
		log.Fatalf("Failed to invoke storage: %v", err)
	}

	// Provide Services injection
	if err := injector.ProvideMulti(services); err != nil {
		log.Fatalf("Failed to invoke service: %v", err)
	}

	// Provide Handler injection
	if err := injector.ProvideMulti(handler); err != nil {
		log.Fatalf("Failed to invoke service: %v", err)
	}

	err = injector.Invoke(func(router *gin.Engine, taskHandler *handlers.TaskHandler) {
		taskHandler.RegisterRoutes(router)
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to invoke dependencies: %v", err)
	}
}
