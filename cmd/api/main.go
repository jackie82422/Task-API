package main

import (
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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
		swagger, err := openapi3.NewLoader().LoadFromFile("./cmd/api/api_doc.yaml")
		if err != nil {
			log.Fatal(err)
		}

		swagger.Servers = nil
		router := gin.Default()
		router.GET("/openapi.json", func(c *gin.Context) {
			c.JSONP(http.StatusOK, swagger)
		})
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi.json")))
		return router
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
