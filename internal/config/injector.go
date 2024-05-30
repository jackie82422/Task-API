package config

import (
	"os"

	"go.uber.org/dig"

	"task-api/internal/domain/task"
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

func (i *Injector) ProvideStorage(inMemoryRepos []interface{}, relationRepos []interface{}) error {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == Mysql {
		dsn := os.Getenv("MYSQL_DSN")
		for _, repoConstructor := range relationRepos {
			if err := i.container.Provide(func() (task.Repository, error) {
				return repoConstructor.(func(string) (task.Repository, error))(dsn)
			}); err != nil {
				return err
			}
		}
	} else {
		for _, repoConstructor := range inMemoryRepos {
			if err := i.container.Provide(repoConstructor.(func() task.Repository)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (i *Injector) ProvideMulti(constructors []interface{}) error {
	for _, constructor := range constructors {
		if err := i.container.Provide(constructor); err != nil {
			return err
		}
	}
	return nil
}

func (i *Injector) Provide(constructor interface{}) error {
	return i.container.Provide(constructor)
}
