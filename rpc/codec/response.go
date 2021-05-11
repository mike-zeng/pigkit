package response

import (
	"pigkit/rpc/codec"
)

type PigResponse struct {
	ServerName string
	MetaData map[string]interface{}
	SerializationType int
	Content interface{}
}

func (res *PigResponse) Bytes() []byte{
	return nil
}

func FrameToPigResponse(frame *codec.Frame) (*PigResponse,error){
	return nil,nil
}