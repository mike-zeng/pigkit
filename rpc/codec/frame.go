package frame

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Frame struct {
	header Header
	body   Body
}

type Header struct {
	Magic uint8    // magic
	Version uint8  // version
	MsgType uint8  // msg type e.g. :   0x0: general req,  0x1: heartbeat
	ReqType uint8  // request type e.g. :   0x0: send and receive,   0x1: send but not receive,  0x2: client stream request, 0x3: server stream request, 0x4: bidirectional streaming request
	CompressType uint8 // compression or not :  0x0: not compression,  0x1: compression
	StreamID uint16    // stream ID
	Length uint32      // total packet length
	Reserved uint32  // 4 bytes reserved
}

type Body struct {
	content []byte
}

func (frame Frame) Bytes()[]byte {
	buffer := bytes.Buffer{}
	// 写入header
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.Magic)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.MsgType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.ReqType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.CompressType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.StreamID)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.Length)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.header.Reserved)
	// 写入body
	_ = binary.Write(&buffer, binary.LittleEndian, frame.body.content)
	return buffer.Bytes()
}

func (frame *Frame) IsReq()bool  {
	return frame.header.MsgType == 0x0
}

func (frame *Frame) IsResp()bool  {
	return frame.header.MsgType == 0x1
}

func BuildFromReader(conn net.Conn)(*Frame,error) {
	return nil,nil
}
