package main

import (
	"awesomeProject/config"
	"awesomeProject/todolist"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func main() {
	container := provide()

	logrus.Info("running rpc server")
	if err := container.Invoke(func(handler todolist.Server) {
		handler.Launch()
	}); err != nil {
		panic(err)
	}

}

func provide() *dig.Container {
	container := dig.New()
	container.Provide(config.NewEnv)
	container.Provide(config.NewDatabase)
	container.Provide(todolist.NewMysqlRepository)
	container.Provide(todolist.NewTodoListServer)

	return container
}
