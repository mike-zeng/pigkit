package server

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
)

// Handler 处理器方法类型
// param1:context
// param2:interface{}，handler实现类
// param3:func(interface{},*codec.PigReq)error, 请求类型转换函数
type Handler func(context.Context,interface{},*codec.PigReq,func(interface{},*codec.PigReq)error)(interface{},error)