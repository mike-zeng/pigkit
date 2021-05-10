package request

import "pigkit/rpc/frame"

type PigReq struct {
	ServerName string
	MetaData map[string]interface{}
	Content interface{}
}

func (req *PigReq) ToFrame() frame.Frame {
	return frame.Frame{
		
	}
}