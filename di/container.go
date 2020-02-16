package di

import (
	"github.com/btomoaki/go-mydi/app"
	"github.com/btomoaki/go-mydi/domain/iface"
	"github.com/btomoaki/go-mydi/domain/service/web"
	"github.com/btomoaki/go-mydi/infra/config"
	"github.com/btomoaki/go-mydi/infra/router"
)

type Di struct {
	Container map[string]interface{}
}

func New() *Di {
	return &Di{
		Container: make(map[string]interface{}),
	}
}

func (di *Di) Env() map[string]string {
	name := "Env"
	if _, ok := di.Container[name]; ok == false {
		item := &config.Env{}

		di.Container[name] = item.LoadEnv()
	}
	return di.Container[name].(map[string]string)
}

func (di *Di) Contoror() *app.Contoror {
	name := "Contoror"
	if _, ok := di.Container[name]; ok == false {
		item := &app.Contoror{EnvMap: di.Env(), Server: di.WebServer(), IndexService: di.IndexService()}
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
