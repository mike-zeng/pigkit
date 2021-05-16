package main

import (
	"github.com/mike-zeng/pigkit/rpc/server"
	"pigkit/example/hello/hello"
)

func main() {

	pigServer := server.NewPigServer(hello.HelloServiceDesc, HelloHandler{})
	options := &server.Options{
		Port: 8181,
	}
	pigServer.Serve(options)
}
