package server

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/config"
	"github.com/mike-zeng/pigkit/rpc/third"
	"github.com/mike-zeng/pigkit/rpc/transport"
)

type Server interface {
	Serve(*Options)
	HandlerReq(context.Context,*codec.PigReq)(*codec.PigResponse,error)
	Close()
}

type PigServer struct {
	desc    *ServiceDesc
	service interface{}
	trans   transport.PigServerTransport
}

func NewPigServer(desc *ServiceDesc, service interface{},confPath string) *PigServer {
	config.InitConf(confPath)
	return &PigServer{desc: desc, service: service}
}

func (p *PigServer) HandlerReq(ctx context.Context,pigReq *codec.PigReq)(*codec.PigResponse,error) {
	serialization := codec.GetSerialization(pigReq.SerializationType)
	handler := p.desc.Methods[pigReq.MethodName]
	respBody, err := handler(ctx, p.service, pigReq, func(req interface{},pigReq *codec.PigReq) error {
		return serialization.Unmarshal(pigReq.Content, req)
	})
	if err != nil {
		return nil, err
	}
	marshal, err := serialization.Marshal(respBody)
	if err != nil {
		return nil, err
	}
	return &codec.PigResponse{
		Content:           marshal,
	},nil
}

func (p *PigServer) Serve(options *Options) {
	// 服务注册
	go p.Register(options)
	p.trans.SetHandlerReq(p.HandlerReq)
	p.trans.ListenAndServer(options.Ip,options.Port)
}

func (p *PigServer) Register(options *Options) {
	// 服务注册
	service := third.GetEtcdService()
	err := service.RegService("123", config.GetConfig().PigServer.ServiceName, "127.0.0.1:8080")
	if err != nil {
		// todo 重试
	}
}


func (p *PigServer) Close() {

}


