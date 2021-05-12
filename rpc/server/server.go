package server

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/transport"
)

type Server interface {
	Serve(*Options)
	HandlerReq(*codec.PigReq)(*codec.PigResponse,error)
	Close()
}

type PigServer struct {
	desc    *ServiceDesc
	service interface{}
	trans   transport.PigServerTransport
}

func NewPigServer(desc *ServiceDesc, service interface{}) *PigServer {
	return &PigServer{desc: desc, service: service}
}

func (p *PigServer) HandlerReq(pigReq *codec.PigReq)(*codec.PigResponse,error) {
	serialization := codec.GetSerialization(pigReq.SerializationType)
	handler := p.desc.Methods[pigReq.MethodName]
	respBody, err := handler(context.Background(), p.service, pigReq, func(req interface{},pigReq *codec.PigReq) error {
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
	p.trans.SetHandlerReq(p.HandlerReq)
	p.trans.ListenAndServer(options.Ip,options.Port)
}

func (p *PigServer) Close() {

}


