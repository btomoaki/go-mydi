package service

import "github.com/gin-gonic/gin"

type WebService struct{}

func (w *WebService) Index() (status int, template string, hundler gin.H) {
	return 200, "index", gin.H{
		"title":   "index",
		"message": "hello world",
	}
}
