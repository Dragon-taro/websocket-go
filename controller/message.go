package controller

import (
	"github.com/gin-gonic/gin"
)

type MessageController struct{}

func (m *MessageController) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"masseage": "hello golang",
	})
}
