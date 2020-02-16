package app

import (
	"github.com/btomoaki/go-mydi/domain/iface"
	"github.com/btomoaki/go-mydi/domain/service/web"
)

type Contoror struct {
	Server       iface.Router
	IndexService web.IndexServiceInterface
}

func (c *Contoror) MapRoute() {
	c.Server.GET("/", c.IndexService.Get())
	//c.Server.PUT("/", c.WebService.Index("put"))
	//c.Server.POST("/", c.WebService.Index("post"))
}

func (c *Contoror) RenderTemplate() {
	c.IndexService.RenderIndex()
}

func (c *Contoror) Run() {
	c.Server.Run(":8080")
}
