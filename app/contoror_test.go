package app_test

import (
	"github.com/btomoaki/go-mydi/app"
	"github.com/btomoaki/go-mydi/domain/iface"
	"github.com/btomoaki/go-mydi/domain/service/web"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Contoror", func() {
	var (
		c                *gomock.Controller
		mockRouter       *iface.MockRouter
		mockIndexService *web.MockIndexServiceInterface
	)
	Context("MapRoute", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRouter = iface.NewMockRouter(c)
			mockIndexService = web.NewMockIndexServiceInterface(c)
		})
		It("Call MapRoute", func() {
			contoror := &app.Contoror{Server: mockRouter, IndexService: mockIndexService}
			mockRouter.EXPECT().GET("/", gomock.Any()).Return(nil)
			mockIndexService.EXPECT().Get().Return(func(*gin.Context) {})
			contoror.MapRoute()
		})
	})
	Context("RenderTemplate", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRouter = iface.NewMockRouter(c)
			mockIndexService = web.NewMockIndexServiceInterface(c)
		})
		It("Call RenderTemplate", func() {
			contoror := &app.Contoror{Server: mockRouter, IndexService: mockIndexService}
			mockIndexService.EXPECT().RenderIndex()
			contoror.RenderTemplate()
		})
	})
	Context("Run", func() {
		BeforeEach(func() {
			c = gomock.NewController(GinkgoT())
			mockRouter = iface.NewMockRouter(c)
		})
		It("Call Run", func() {
			contoror := &app.Contoror{Server: mockRouter, IndexService: nil}
			mockRouter.EXPECT().Run(":8080").Return(nil)
			contoror.Run()
		})
	})

})
