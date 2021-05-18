package third

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_tracerClient_StartSpan(t *testing.T) {
	client := GetTracerClient()
	ctx := context.Background()
	ctx,span := client.StartSpanClient(ctx, "test2")
	time.Sleep(1*time.Second)
	span.Finish()
	time.Sleep(10*time.Second)
	fmt.Println(ctx)
}
