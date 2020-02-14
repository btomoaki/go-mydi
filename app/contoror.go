package app

import (
	"../domain/iface"
	"../domain/service"
	"github.com/gin-gonic/gin"
)

type Contoror struct {
	Server     iface.Router
	Render     iface.Render
	WebService *service.WebService
}

func (c *Contoror) MapRoute() {
	c.Server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(c.WebService.Index())
	})

}

func (c *Contoror) RenderTemplate() {

	c.Render.AddFromFiles("index",
		"templates/common/base.html",
		"templates/common/header.html",
		"templates/index.html")

}

func (c *Contoror) Run() {
	c.Server.Run(":8080")
}
