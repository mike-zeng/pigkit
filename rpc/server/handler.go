package server

import (
	"context"
	"pigkit/rpc/codec"
)

type Handler func(context.Context,*codec.PigReq,func(interface{})error)(*codec.PigResponse,error)