package main

import (
	"github.com/Dragon-taro/websocket-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	message := controller.MessageController{}
	r.GET("/", message.Root)
	r.Run()
}
