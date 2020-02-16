package web_test

import (
	reflect "reflect"

	"github.com/btomoaki/go-mydi/domain/iface"
	"github.com/btomoaki/go-mydi/service/web"
	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IndexService", func() {
	var (
		c          *gomock.Controller
		mockRender *iface.MockRender
	)

	Context("Get", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRender = iface.NewMockRender(c)
		})
		It("Call Get", func() {
			indexService := &web.IndexService{Render: mockRender}
			get := indexService.Get()
			Expect(reflect.ValueOf(get).String()).To(Equal(reflect.ValueOf(func(ctx *gin.Context) {
				ctx.HTML(200, "index", gin.H{
					"title":   "index",
					"message": "hello world",
				})
			}).String()))
		})
	})

	Context("NotAllowed", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRender = iface.NewMockRender(c)
		})
		It("Call NotAllowed", func() {
			indexService := &web.IndexService{Render: mockRender}
			get := indexService.NotAllowed()
			Expect(reflect.ValueOf(get).String()).To(Equal(reflect.ValueOf(func(ctx *gin.Context) {
				ctx.HTML(405, "NotAllowed", gin.H{
					"title":   "Method Not Allowed.",
					"message": "Sorry... Method Not Allowed.",
				})
			}).String()))
		})
	})

	Context("RenderIndex", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRender = iface.NewMockRender(c)
		})
		It("Call RenderIndex", func() {
			indexService := &web.IndexService{Render: mockRender}
			mockRender.EXPECT().AddFromFiles("index",
				"templates/common/base.html",
				"templates/common/header.html",
				"templates/index.html")
			mockRender.EXPECT().AddFromFiles("NotAllowed",
				"templates/common/base.html",
				"templates/common/header.html",
				"templates/index.html")
			indexService.RenderIndex()
		})
	})

})
