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
	respBody, err := handler(context.Background(), handler, pigReq, func(req interface{},pigReq *codec.PigReq) error {
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
	p.trans.ListenAndServer(options.Ip,options.Port)
}

func (p *PigServer) Close() {

}


