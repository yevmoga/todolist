package todolist

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (t *TodoListServer) put() func(c *gin.Context) {
	return func(c *gin.Context) {

		var task Task

		if err := c.BindJSON(&task); err != nil {
			logrus.WithError(err).Error("error parsing form")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := task.ValidateByFields([]string{validateById, validateByDone}); err != nil {
			logrus.WithError(err).Error("validation error")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := t.repository.UpdateTask(&task); err != nil {
			logrus.WithError(err).Error("error creating new daily task")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, task)
	}
}
