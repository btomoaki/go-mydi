package router

import (
	"html/template"

	multitemplate "github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin/render"
)

type Render struct {
	Template multitemplate.Render
}

func NewRender() *Render {
	render := new(Render)
	render.Template = multitemplate.New()
	return render
}

func (r *Render) Instance(name string, data interface{}) render.Render {
	return r.Template.Instance(name, data)
}
func (r *Render) AddFromFiles(name string, files ...string) *template.Template {
	return r.Template.AddFromFiles(name, files...)
}
