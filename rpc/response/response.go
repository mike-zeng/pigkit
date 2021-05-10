package response

import "pigkit/rpc/frame"

type PigResponse struct {

}

func (res *PigResponse) ToFrame() frame.Frame {
	return frame.Frame{

	}
}

func FrameToPigResponse(frame *frame.Frame) (*PigResponse,error){
	return nil,nil
}