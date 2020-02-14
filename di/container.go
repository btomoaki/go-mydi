package di

import (
	"../app"
	"../domain/iface"
	"../domain/service"
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
		item := &app.Contoror{Server: di.WebServer(), Render: di.Render(), WebService: di.WebService()}
		di.Container[name] = item
	}
	return di.Container[name].(*app.Contoror)
}

func (di *Di) WebService() *service.WebService {
	name := "WebService"
	if _, ok := di.Container[name]; ok == false {
		item := &service.WebService{}
		di.Container[name] = item
	}
	return di.Container[name].(*service.WebService)
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
