
package main

import (
	"github.com/mike-zeng/pigkit/rpc/server"
	desc "example/test/gen/example/test"
)

func main() {
	pigServer := server.NewPigServer(desc.TestServiceDesc, TestHandler{}, "./pig.yaml")
	options := &server.Options{
		Port: 8080,
	}
	pigServer.Serve(options)
}