package transport

import (
	"fmt"
	"net"
	"pigkit/rpc/codec"
	"strconv"
)

type ServerTransport interface {
	SetHandlerReq(func(*codec.PigReq)*codec.PigResponse,error)
	ListenAndServer(ip string,port int)
}


type PigServerTransport struct {
	handlerReq func(*codec.PigReq)(*codec.PigResponse, error)
}

func (trans *PigServerTransport) SetHandlerReq(handlerReq func(*codec.PigReq) (*codec.PigResponse, error)) {
	trans.handlerReq = handlerReq
}

func (trans *PigServerTransport) ListenAndServer(ip string,port int)  {
	address := ip + ":"+ strconv.Itoa(port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	for {
		conn , err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go trans.handleConn(conn)
	}
}

func (trans *PigServerTransport)handleConn(conn net.Conn)  {
	frame, err := ReadFrame(conn)
	if err != nil {
		// todo
		return
	}
	// 解码req
	request, err := codec.EnCodingRequest(frame)
	if err != nil {
		return
	}
	resp,err := trans.handlerReq(request)
	if err != nil {
		// todo 处理错误
	}
	respFrame := codec.CodingResponseToFrame(resp)
	_, err = conn.Write(respFrame.Bytes())
	if err != nil {
		fmt.Println("写入失败")
	}
}