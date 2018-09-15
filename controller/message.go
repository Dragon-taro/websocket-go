package controller

import (
	"fmt"

	"github.com/Dragon-taro/websocket-go/model"
	"github.com/gin-gonic/gin"
)

type MessageController struct{}

func (m *MessageController) Get(c *gin.Context) {
	messages, err := model.GetAll()
	if err != nil {
		fmt.Println(err)
		// return err
		c.JSON(500, err)
		return
	}
	c.JSON(200, messages)
}
