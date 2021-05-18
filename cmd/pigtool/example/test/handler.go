
package main

import (
	"context"
	desc "example/test/gen/example/test"
)

type TestHandler struct{
	// handler struct
}

// Test pig tool auto gen
func (handler *TestHandler)Test(ctx context.Context, req *desc.TestRequest)(*desc.TestResponse, error) {
	// todo impl handler method
	return nil, nil
}