package codec

import (
	"pigkit/rpc/env"
)

func CodingRequestToFrame(req *PigReq) *Frame {
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


func CodingResponseToFrame(resp *PigResponse) *Frame {
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