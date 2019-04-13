package todolist

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Launch()
}

func NewTodoListServer(repository Repository) Server {
	return &TodoListServer{
		repository: repository,
	}
}

type TodoListServer struct {
	repository Repository
}

func (t *TodoListServer) Launch() {
	router := gin.Default()

	router.GET("/", t.get())
	router.POST("/", t.post())
	router.PUT("/", t.put())
	router.DELETE("/", t.delete())

	if err := router.Run(); err != nil {
		panic(err)
	}
}
