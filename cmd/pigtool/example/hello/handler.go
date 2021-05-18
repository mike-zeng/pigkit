
package main

import (
	"context"
	desc "example/hello/gen/example/hello"
)

type HelloHandler struct{
	// handler struct
}

// Hello pig tool auto gen
func (handler *HelloHandler)Hello(ctx context.Context, req *desc.PigRequest)(*desc.PigResponse, error) {
	// todo impl handler method
	return nil, nil
}