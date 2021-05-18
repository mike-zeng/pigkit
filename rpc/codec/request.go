package codec

import (
	"encoding/json"
)

type PigReq struct {
	ServiceName       string
	MethodName        string
	MetaData          Metadata
	SerializationType string
	Content           []byte
}

func (req *PigReq) Bytes() []byte {
	marshal, _ := json.Marshal(req)
	return marshal
}