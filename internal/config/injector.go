package config

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"os"
	"task-api/internal/domain/task"
	"task-api/internal/handlers"
	"task-api/internal/infrastructure/persistence/memory"
	"task-api/internal/infrastructure/persistence/mysql"
)

const (
	Mysql = "mysql"
)

type Injector struct {
	container *dig.Container
}

func NewInjector() *Injector {
	return &Injector{
		container: dig.New(),
	}
}

func (i *Injector) Invoke(function interface{}) error {
	return i.container.Invoke(function)
}

func (i *Injector) InvokeStorage() error {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == Mysql {
		dsn := os.Getenv("MYSQL_DSN")
		err := i.container.Provide(func() (task.Repository, error) {
			return mysql.NewMySQLTaskRepository(dsn)
		})
		if err != nil {
			return err
		}
	} else {
		err := i.container.Provide(func() task.Repository {
			return memory.NewInMemoryTaskRepository()
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Injector) RegisterHandlers(router *gin.Engine, constructors []interface{}) error {
	for _, constructor := range constructors {
		if err := i.container.Invoke(constructor); err != nil {
			return err
		}
		handler := constructor.(handlers.Handler)
		handler.RegisterRoutes(router)
	}
	return nil
}
