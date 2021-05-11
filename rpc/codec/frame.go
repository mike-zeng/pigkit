package codec

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Frame struct {
	Header *Header
	Body   *Body
}

const FrameHeadLen = 15
type Header struct {
	Magic uint8    // magic
	Version uint8  // version
	MsgType uint8  // msg type e.g. :   0x0: general req, 0x1: general resp,  0x2: heartbeat
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
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.Magic)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.Version)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.MsgType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.ReqType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.CompressType)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.StreamID)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.Length)
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Header.Reserved)
	// 写入body
	_ = binary.Write(&buffer, binary.LittleEndian, frame.Body.content)
	return buffer.Bytes()
}

func (frame *Frame) IsReq()bool  {
	return frame.Header.MsgType == 0x0
}

func (frame *Frame) IsResp()bool  {
	return frame.Header.MsgType == 0x1
}

func BuildFromReader(conn net.Conn)(*Frame,error) {
	return nil,nil
}
