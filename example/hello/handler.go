package main

import (
	"context"
	"fmt"
	"pigkit/example/hello/hello"
)

type HelloHandler struct {

}

func (h HelloHandler) Hello(ctx context.Context, req *hello.Request) (*hello.Response, error) {
	fmt.Println(req.Name)
	resp := &hello.Response{}
	resp.Msg = "hello"+req.Name
	return resp,nil
}