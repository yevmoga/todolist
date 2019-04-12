package todolist

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Launch()
}

func NewTodoListServer() Server {
	return &TodoListServer{}
}

type TodoListServer struct{}

func (TodoListServer) Launch() {
	router := gin.Default()

	router.GET("/", GetTodoList)
	router.POST("/", PostTodoList)
	router.PUT("/", PutTodoList)
	router.DELETE("/", DeleteTodoList)

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func GetTodoList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostTodoList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Hello World",
	})
}

func PutTodoList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Put Hello World",
	})
}

func DeleteTodoList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Hello World",
	})
}
