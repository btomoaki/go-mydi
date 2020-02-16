package main

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/btomoaki/go-mydi/di"
)

func main() {
	Di := di.New()
	wg := &sync.WaitGroup{}
	contoror := Di.Contoror()
	contoror.MapRoute()
	contoror.RenderTemplate()
	wg.Add(1)
	go func() {
		defer wg.Done()

		contoror.Run()
	}()
	ioutil.WriteFile("/tmp/app-initialized", []byte("\n"), os.ModePerm)
	wg.Wait()
}
