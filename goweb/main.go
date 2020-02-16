package main

import (
	"io/ioutil"
	"os"

	"github.com/btomoaki/go-mydi/di"
)

func main() {
	Di := di.New()
	contoror := Di.Contoror()
	contoror.MapRoute()
	contoror.RenderTemplate()
	contoror.Run()
	ioutil.WriteFile("/tmp/app-initialized", []byte("\n"), os.ModePerm)
}
