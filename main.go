package main

import (
	"html/template"
	"net/http"

	"github.com/Dragon-taro/websocket-go/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	message := controller.MessageController{}

	// html
	r.GET("/", func(c *gin.Context) {
		// テンプレート設定
		html := template.Must(template.ParseFiles("templates/base.tmpl", "templates/index.tmpl"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.tmpl", gin.H{})
	})

	// api
	r.GET("/messages", message.Root)

	// static
	r.Static("/public", "./public")

	r.Run()
}
