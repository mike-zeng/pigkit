package server

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
)

type Handler func(context.Context,interface{},*codec.PigReq,func(interface{},*codec.PigReq)error)(interface{},error)