package main

import (
	"./di"
)

func main() {
	Di := di.New()
	contoror := Di.Contoror()
	contoror.MapRoute()
	contoror.RenderTemplate()
	contoror.Run()
}
