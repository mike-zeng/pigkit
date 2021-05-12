package client

import (
	"context"
	"fmt"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/pool"
	"github.com/mike-zeng/pigkit/rpc/transport"
	"sync"
	"testing"
	"time"
)

type TestReq struct {
	A int `json:"a"`
	B int `json:"b"`
}

type TestResp struct {
	Sum int `json:"sum"`
}
func TestPigClient_Call(t *testing.T) {
	manage := pool.NewManage(pool.DefaultLoadBalancer())
	manage.RegisterPool(pool.ConnInfo{
		ServerName:  "Calculation",
		IpAddr:      "127.0.0.1",
		Port:        8181,
		MaxConnNum:  10,
		MinConnNum:  10,
	})
	clientTransport := transport.NewPigClientTransport(manage)
	client := NewPigClient(clientTransport)
	req := &TestReq{
		A: 1,
		B: 2,
	}
	options := Options{
		serviceName:       "Calculation",
		method:            "add",
		SerializationType: "json",
	}
	group := sync.WaitGroup{}
	group.Add(10000)
	t1 := time.Now()
	for i:=0;i<10000;i++{
		go func() {
			resp, err := client.Call(context.Background(), req, options)
			if err != nil {
				// todo
				return
			}
			serialization := codec.GetSerialization("json")
			testResp := TestResp{}
			err = serialization.Unmarshal(resp, &testResp)
			if err != nil {
				return
			}
			group.Done()
		}()
	}
	group.Wait()
	t2 := time.Now()
	fmt.Println(t2.UnixNano()-t1.UnixNano())
}
