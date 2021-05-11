package transport

import (
	"fmt"
	"pigkit/rpc/codec"
	"pigkit/rpc/pool"
	"testing"
)

func TestPigClientTransport_SyncSend(t *testing.T) {
	manage := pool.NewManage(pool.LoadBalancer{})
	manage.RegisterPool(pool.ConnInfo{
		ServerName: "hello",
		IpAddr:     "127.0.0.1",
		Port:       8181,
		MaxConnNum:  10,
		MinConnNum:  10,
		MaxIdleTime: 2000,
	})
	transport := NewPigClientTransport(manage)
	req := codec.PigReq{
		ServerName:        "hello",
		MetaData:          nil,
		SerializationType: 0,
		Content:           "hello world",
	}
	frame := codec.CodingRequestToFrame(&req)
	send, err := transport.SyncSend("hello", frame.Bytes())
	if err != nil {
		fmt.Println("发送失败")
		return
	}
	response, err := codec.EnCodingResponse(send)
	if err != nil {
		fmt.Println("发送失败")
		return
	}
	fmt.Println(response)
}
