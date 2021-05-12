package server

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"testing"
)


type TestReq struct {
	A int `json:"a"`
	B int `json:"b"`
}

type TestResp struct {
	Sum int `json:"sum"`
}

type CalculationService interface {
	Add(ctx context.Context,req *TestReq)(*TestResp,error)
}


type CalculationServiceImpl struct {
}

func (impl CalculationServiceImpl) Add(ctx context.Context,req *TestReq)(*TestResp,error) {
	resp := &TestResp{}
	resp.Sum = req.A +req.B
	return resp,nil
}

func GenAddHandler(ctx context.Context,s interface{},pigReq *codec.PigReq,f func(interface{},*codec.PigReq)error)(interface{},error)  {
	req := TestReq{}
	err := f(&req, pigReq)
	if err != nil {
		return nil, err
	}
	impl := s.(CalculationServiceImpl)
	return impl.Add(ctx,&req)
}

func TestNewPigServer(t *testing.T) {
	desc := &ServiceDesc{
		ServiceName: "Calculation",
		Methods: map[string]Handler{
			"add": GenAddHandler,
		},
	}
	server := NewPigServer(desc, CalculationServiceImpl{})
	options := &Options{
		Port: 8181,
	}
	server.Serve(options)
}
