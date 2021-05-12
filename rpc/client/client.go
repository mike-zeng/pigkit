package client

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/transport"
)

type Client interface {
	Call(ctx context.Context,req interface{},options Options)([]byte,error)
}

type PigClient struct {
	trans transport.ClientTransport
}

func NewPigClient(trans transport.ClientTransport) *PigClient {
	return &PigClient{trans: trans}
}

func (p *PigClient) Call(ctx context.Context, req interface{}, options Options) ([]byte, error) {
	serialization := codec.GetSerialization(options.SerializationType)
	contentData, err := serialization.Marshal(req)
	pigReq := &codec.PigReq{
		ServiceName: options.serviceName,
		MethodName:  options.method,
		MetaData:    nil,
		Content:     contentData,
		SerializationType: options.SerializationType,
	}
	frame := codec.DecodingRequestToFrame(pigReq)
	if err != nil {
		return nil, err
	}
	send, err := p.trans.SyncSend(options.serviceName, frame.Bytes())
	if err != nil {
		return nil, err
	}
	response, err := codec.EncodingFrameToResponse(send)
	if err != nil {
		return nil, err
	}
	return response.Content,nil
}