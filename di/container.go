package di

import (
	"../app"
	"../domain/iface"
	"../domain/service/web"
	"../infra/router"
)

type Di struct {
	Container map[string]interface{}
}

func New() *Di {
	return &Di{
		Container: make(map[string]interface{}),
	}
}

func (di *Di) Contoror() *app.Contoror {
	name := "Contoror"
	if _, ok := di.Container[name]; ok == false {
		item := &app.Contoror{Server: di.WebServer(), IndexService: di.IndexService()}
		di.Container[name] = item
	}
	return di.Container[name].(*app.Contoror)
}

func (di *Di) IndexService() *web.IndexService {
	name := "IndexService"
	if _, ok := di.Container[name]; ok == false {
		item := &web.IndexService{Render: di.Render()}
		di.Container[name] = item
	}
	return di.Container[name].(*web.IndexService)
}

func (di *Di) WebServer() iface.Router {
	name := "WebServer"
	if _, ok := di.Container[name]; ok == false {
		item := router.NewRouter(di.Render())
		di.Container[name] = item
	}
	return di.Container[name].(iface.Router)
}

func (di *Di) Render() iface.Render {
	name := "Render"
	if _, ok := di.Container[name]; ok == false {
		item := router.NewRender()
		di.Container[name] = item
	}
	return di.Container[name].(iface.Render)
}
