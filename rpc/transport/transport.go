package transport

import (
	"bytes"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"io"
	"net"
)

func ReadFrame(conn net.Conn) (*codec.Frame,error){
	// read frame header
	frameHeaderData := make([]byte,codec.FrameHeadLen)
	if num, err := io.ReadFull(conn, frameHeaderData); num != codec.FrameHeadLen || err != nil {
		return nil,err
	}
	frameHeader, err := codec.ReadFrameHeader(bytes.NewBuffer(frameHeaderData))
	if err != nil {
		return nil, err
	}
	// read frame body
	frameBodyData := make([]byte,frameHeader.Length)
	if num, err := io.ReadFull(conn, frameBodyData); num != int(frameHeader.Length) || err != nil {
		return nil,err
	}
	frameBody, err := codec.ReadFrameBody(bytes.NewBuffer(frameBodyData))
	if err != nil {
		return nil, err
	}
	return &codec.Frame{
		Header: frameHeader,
		Body: frameBody,
	},nil
}
