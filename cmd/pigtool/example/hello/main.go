
package main

import (
	"github.com/mike-zeng/pigkit/rpc/server"
	desc "example/hello/gen/example/hello"
)

func main() {
	pigServer := server.NewPigServer(desc.HelloServiceDesc, HelloHandler{}, "./pig.yaml")
	options := &server.Options{
		Port: 8080,
	}
	pigServer.Serve(options)
}