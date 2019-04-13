package todolist

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (t *TodoListServer) delete() func(c *gin.Context) {
	return func(c *gin.Context) {

		var task Task

		if err := c.BindJSON(&task); err != nil {
			logrus.WithError(err).Error("error parsing form")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := task.ValidateByFields([]string{validateById}); err != nil {
			logrus.WithError(err).Error("validation error")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := t.repository.DeleteTask(&task); err != nil {
			logrus.WithError(err).Error("error deleting daily task")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, map[string]string{"result": "task deleted"})
	}
}
