package transport

import (
	"errors"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/pool"
)

type ClientTransport interface {
	SyncSend(serverName string,bytes []byte)(*codec.Frame,error)
}


type PigClientTransport struct {
	manage *pool.Manage
}

func NewPigClientTransport(manage *pool.Manage) *PigClientTransport {
	return &PigClientTransport{manage: manage}
}

func (trans *PigClientTransport) SyncSend(serverName string, data []byte)(*codec.Frame,error) {
	netPool := trans.manage.GetPool(serverName)
	if netPool == nil{
		 return nil,errors.New("server not exist")
	}
	node, err := netPool.Get()
	defer func() {
		netPool.Return(node)
	}()
	if err != nil||node==nil {
		return nil,errors.New("no free node")
	}
	// write data
	_, err = node.Conn.Write(data)
	if err != nil{
		return nil,err
	}
	// read frame from conn
	frame, err := ReadFrame(node.Conn)
	if err != nil {
		return nil, err
	}
	return frame,nil
}

