package iface

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Run(port ...string) (err error)
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}
