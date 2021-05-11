package codec

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"pigkit/rpc/env"
)


func EnCodingResponse(frame *Frame) (*PigResponse,error) {

	if !frame.IsResp() {
		return nil,errors.New("frame msg type error,need resp type")
	}
	content := frame.Body.content
	pigResponse := PigResponse{}
	err := json.Unmarshal(content, &pigResponse)
	return &pigResponse,err
}

func EnCodingRequest(frame *Frame) (*PigReq,error) {

	if !frame.IsReq() {
		return nil,errors.New("frame msg type error,need req type")
	}
	content := frame.Body.content
	pigReq := PigReq{}
	err := json.Unmarshal(content, &pigReq)
	return &pigReq,err
}

func ReadFrameHeader(buffer *bytes.Buffer) (*Header,error)  {
	frameHeader := Header{}
	var(
		magic uint8
		version uint8
		msgType uint8
		reqType uint8
		compressType uint8
		streamId uint16
		length uint32
		reserved uint32
	)
	err := binary.Read(buffer, binary.LittleEndian, &magic)
	if err != nil {
		return nil,err
	}
	// magic num check
	if env.CheckMagicNum(magic){
		frameHeader.Magic = magic
	}else {
		return nil,errors.New("magic num error")
	}

	err = binary.Read(buffer, binary.LittleEndian, &version)
	if err != nil {
		return nil,err
	}
	// check version
	if env.CheckVersion(version) {
		frameHeader.Version = version
	}else {
		return nil,errors.New("version num error")
	}
	// check msgType
	err = binary.Read(buffer, binary.LittleEndian, &msgType)
	if err != nil {
		return nil,err
	}
	if env.CheckMsgType(msgType) {
		frameHeader.MsgType = msgType
	}else {
		return nil,errors.New("msgType error")
	}
	// check reqType
	err = binary.Read(buffer, binary.LittleEndian, &reqType)
	if err != nil {
		return nil,err
	}
	if env.CheckReqType(reqType) {
		frameHeader.ReqType = reqType
	}else {
		return nil,errors.New("reqType error")
	}

	// check compress type
	err = binary.Read(buffer, binary.LittleEndian, &compressType)
	if err != nil {
		return nil,err
	}
	if compressType == 0x0 {
		frameHeader.CompressType = compressType
	}else {
		// 暂时不支持压缩
		return nil,errors.New("compress type error")
	}

	// check streamId
	err = binary.Read(buffer, binary.LittleEndian, &streamId)
	if err != nil {
		return nil,err
	}
	frameHeader.StreamID = streamId

	// read length
	err = binary.Read(buffer, binary.LittleEndian, &length)
	if err != nil {
		return nil,err
	}
	frameHeader.Length = length

	// read Reserved
	err = binary.Read(buffer, binary.LittleEndian, &reserved)
	if err != nil {
		return nil,err
	}
	frameHeader.Reserved = reserved
	return &frameHeader,nil
}

func ReadFrameBody(buffer *bytes.Buffer) (*Body,error){
	// read Body
	frameBody := Body{}
	var bodyBytes = make([]byte,buffer.Len())
	err := binary.Read(buffer,binary.LittleEndian,bodyBytes)
	if err != nil {
		return nil,err
	}
	frameBody.content = bodyBytes[:]
	return &frameBody,nil
}