package server

import (
	"context"
	"pigkit/rpc/codec"
	"pigkit/rpc/transport"
)

type Server interface {
	Serve(*Options)
	HandlerReq(*codec.PigReq)(*codec.PigResponse,error)
	Close()
}

type PigServer struct {
	desc ServiceDesc
	handler interface{}
	trans transport.ServerTransport
}

func NewPigServer(desc ServiceDesc, handler interface{}) *PigServer {
	return &PigServer{desc: desc, handler: handler}
}

func (p *PigServer) HandlerReq(pigReq *codec.PigReq)(*codec.PigResponse,error) {
	serialization := &codec.PbSerialization{}
	handler := p.desc.Methods[pigReq.MethodName]
	return handler(context.Background(), pigReq, func(req interface{}) error {
		return serialization.Unmarshal(pigReq.Content, req)
	})
}

func (p *PigServer) Serve(options *Options) {
	p.trans.ListenAndServer(options.Ip,options.Port)
}

func (p *PigServer) Close() {

}


