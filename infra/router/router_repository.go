package router

import (
	gin "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Router struct {
	Gin *gin.Engine
}

func NewRouter(render render.HTMLRender) *Router {
	router := new(Router)
	router.Gin = gin.Default()
	router.Gin.HTMLRender = render
	return router
}

func (r *Router) Run(addr ...string) (err error) {
	return r.Gin.Run(addr...)
}

func (r *Router) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Gin.GET(relativePath, handlers...)
}
