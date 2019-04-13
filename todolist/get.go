package todolist

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (t *TodoListServer) get() func(c *gin.Context) {
	return func(c *gin.Context) {
		list, err := t.repository.GetAllForToday()
		if err != nil {
			logrus.WithError(err).Error("error getting daily list")
			c.JSON(200, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, list)
	}
}
