package transport

import (
	"errors"
	"pigkit/rpc/frame"
	"pigkit/rpc/pool"
	"pigkit/rpc/request"
	"pigkit/rpc/response"
)

type ClientTransport interface {
	SyncSend(req *request.PigReq)(*response.PigResponse,error)
}


type PigClientTransport struct {
	manage pool.Manage
}

func (trans *PigClientTransport) SyncSend(req *request.PigReq)(*response.PigResponse,error) {
	bytes := req.ToFrame().Bytes()
	netPool := trans.manage.GetPool(req.ServerName)
	if netPool == nil{
		 return nil,errors.New("server not exist")
	}
	node, err := netPool.Get()
	if err != nil||node==nil {
		return nil,errors.New("no free node")
	}
	_, err = node.Conn.Write(bytes)
	if err != nil{
		return nil,errors.New("write error")
	}
	// get response
	f,err := frame.BuildFromReader(node.Conn)
	if err != nil || f ==nil{
		// todo
	}
	return response.FrameToPigResponse(f)
}