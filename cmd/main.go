package main

import (
	"github.com/btomoaki/go-mydi/di"
)

func main() {
	Di := di.New()
	contoror := Di.Contoror()
	contoror.MapRoute()
	contoror.RenderTemplate()
	contoror.Run()
}
