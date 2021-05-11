package server

import (
	"context"
	"pigkit/rpc/codec"
)

type Handler func(context.Context,interface{},interface{},func(interface{},*codec.PigReq)error)(interface{},error)