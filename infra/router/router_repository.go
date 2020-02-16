package router

import (
	"fmt"
	"io/ioutil"
	"os"

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
	runerr := r.Gin.RunUnix("/tmp/nginx.socket")
	if runerr == nil {
		ioutil.WriteFile("/tmp/app-initialized", []byte("\n"), os.ModePerm)
	}
	stat, err := os.Stat("/tmp/app-initialized")
	fmt.Println(fmt.Sprintf("%#v,%#v", stat, err))
	return runerr
}

func (r *Router) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Gin.GET(relativePath, handlers...)
}
func (r *Router) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Gin.PUT(relativePath, handlers...)
}
func (r *Router) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Gin.POST(relativePath, handlers...)
}
