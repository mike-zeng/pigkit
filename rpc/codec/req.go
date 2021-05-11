package codec

import "encoding/json"

type PigReq struct {
	ServerName string
	MethodName string
	MetaData map[string]interface{}
	SerializationType int
	Content []byte
}

func (req *PigReq) Bytes() []byte {
	marshal, _ := json.Marshal(req)
	return marshal
}