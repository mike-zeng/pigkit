package client

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/pool"
	"github.com/mike-zeng/pigkit/rpc/transport"
	"sync"
)

type Client interface {
	Call(ctx context.Context,req interface{},options Options)([]byte,error)
	Register(proxy Proxy)
}

type pigClient struct {
	trans transport.ClientTransport
}

func (p *pigClient) Register(proxy Proxy) {
	p.trans.Manage().RegisterService(proxy.GetProxyServiceName())
	proxy.RegisterCallback(p)
}

// 获取pigClient 单例模式
var ins *pigClient
var once sync.Once

func GetDefaultClient()Client {
	once.Do(func() {
		manage := pool.NewManage(pool.DefaultLoadBalancer())
		clientTransport := transport.NewPigClientTransport(manage)
		ins = &pigClient{trans: clientTransport}
	})
	return ins
}

func (p *pigClient) Call(ctx context.Context, req interface{}, options Options) ([]byte, error) {
	serialization := codec.GetSerialization(options.SerializationType)
	contentData, err := serialization.Marshal(req)
	if err != nil {
		return nil,err
	}
	pigReq := &codec.PigReq{
		ServiceName: options.ServiceName,
		MethodName:  options.Method,
		MetaData:    nil,
		Content:     contentData,
		SerializationType: options.SerializationType,
	}
	frame := codec.DecodingRequestToFrame(pigReq)
	if err != nil {
		return nil, err
	}
	send, err := p.trans.SyncSend(options.ServiceName, frame.Bytes())
	if err != nil {
		return nil, err
	}
	response, err := codec.EncodingFrameToResponse(send)
	if err != nil {
		return nil, err
	}
	return response.Content,nil
}

