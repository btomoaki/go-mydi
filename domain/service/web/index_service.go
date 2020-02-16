//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package web

import (
	"github.com/btomoaki/go-mydi/domain/iface"

	"github.com/gin-gonic/gin"
)

type IndexServiceInterface interface {
	Get() func(*gin.Context)
	NotAllowed() func(*gin.Context)
	RenderIndex()
}

type IndexService struct {
	Render iface.Render
}

func (s *IndexService) Get() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(200, "index", gin.H{
			"title":   "index",
			"message": "hello world",
		})
	}
}

func (s *IndexService) NotAllowed() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(405, "NotAllowed", gin.H{
			"title":   "Method Not Allowed.",
			"message": "Sorry... Method Not Allowed.",
		})
	}
}

func (s *IndexService) RenderIndex() {
	s.Render.AddFromFiles("index",
		"templates/common/base.html",
		"templates/common/header.html",
		"templates/index.html")
	s.Render.AddFromFiles("NotAllowed",
		"templates/common/base.html",
		"templates/common/header.html",
		"templates/index.html")
}
