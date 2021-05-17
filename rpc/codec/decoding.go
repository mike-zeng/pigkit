package codec

import (
	"github.com/mike-zeng/pigkit/rpc/env"
)

// DecodingRequestToFrame 将Request对象封装成帧
func DecodingRequestToFrame(req *PigReq) *Frame {
	frame := Frame{
		Header: &Header{
			Magic:        env.MagicNum,
			Version:      env.Version,
			MsgType:      env.MsgReqType,
			ReqType:      env.ReqTypeSendAndRec,
		},
		Body: &Body{content: req.Bytes()},
	}
	frame.Header.Length = uint32(len(req.Bytes()))
	return &frame
}

// DecodingResponseToFrame 将Response封装成帧
func DecodingResponseToFrame(resp *PigResponse) *Frame {
	frame := Frame{
		Header: &Header{
			Magic:        env.MagicNum,
			Version:      env.Version,
			MsgType:      env.MsgRespType,
		},
		Body: &Body{content: resp.Bytes()},
	}
	frame.Header.Length = uint32(len(resp.Bytes()))
	return &frame
}