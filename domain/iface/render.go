package iface

import (
	"html/template"

	"github.com/gin-gonic/gin/render"
)

type Render interface {
	AddFromFiles(string, ...string) *template.Template
	Instance(string, interface{}) render.Render
}
