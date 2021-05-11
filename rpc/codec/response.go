package codec

import "encoding/json"

type PigResponse struct {
	ServerName string
	MetaData map[string]interface{}
	SerializationType int
	Content interface{}
}

func (res *PigResponse) Bytes() []byte{
	marshal, _ := json.Marshal(res)
	return marshal
}
