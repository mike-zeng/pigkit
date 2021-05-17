package codec

import (
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"math"
	"sync"
)

// Serialization 序列化
type Serialization interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

const (
	PbType = "proto"
	JsonType = "json"
)

var pbSerialization PbSerialization
var jsonSerialization JsonSerialization

// GetSerialization 获取序列化器
func GetSerialization(typeStr string)Serialization  {
	switch typeStr {
	case PbType:
		return &pbSerialization
	case JsonType:
		return &jsonSerialization
	default:
		return nil
	}
}

// PbSerialization 基于proto实现的序列化器
type PbSerialization struct{

}

func (d *PbSerialization) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("marshal nil interface{}")
	}
	if pm, ok := v.(proto.Marshaler); ok {
		return pm.Marshal()
	}
	buffer := bufferPool.Get().(*cachedBuffer)
	protoMsg := v.(proto.Message)
	lastMarshaledSize := make([]byte, 0, buffer.lastMarshaledSize)
	buffer.SetBuf(lastMarshaledSize)
	buffer.Reset()

	if err := buffer.Marshal(protoMsg); err != nil {
		return nil, err
	}
	data := buffer.Bytes()
	buffer.lastMarshaledSize = upperLimit(len(data))
	buffer.SetBuf(nil)
	bufferPool.Put(buffer)
	return data, nil
}

func (d *PbSerialization) Unmarshal(data []byte, v interface{}) error {
	if data == nil || len(data) == 0 {
		return errors.New("unmarshal nil or empty bytes")
	}
	protoMsg := v.(proto.Message)
	protoMsg.Reset()

	if pu, ok := protoMsg.(proto.Unmarshaler); ok {
		// 可以 unmarshal 自身，无需 buffer
		return pu.Unmarshal(data)
	}

	buffer := bufferPool.Get().(*cachedBuffer)
	buffer.SetBuf(data)
	err := buffer.Unmarshal(protoMsg)
	buffer.SetBuf(nil)
	bufferPool.Put(buffer)
	return err
}

var bufferPool = &sync.Pool{
	New : func() interface {} {
		return &cachedBuffer {
			Buffer : proto.Buffer{},
			lastMarshaledSize : 16,
		}
	},
}

type cachedBuffer struct {
	proto.Buffer
	lastMarshaledSize uint32
}

func upperLimit(val int) uint32 {
	if val > math.MaxInt32 {
		return uint32(math.MaxInt32)
	}
	return uint32(val)
}

// JsonSerialization 基于json实现的序列化
type JsonSerialization struct{

}

func (j *JsonSerialization) Marshal(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func (j *JsonSerialization) Unmarshal(bytes []byte, i interface{}) error {
	return json.Unmarshal(bytes, i)
}
