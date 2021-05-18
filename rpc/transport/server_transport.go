package transport

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	"github.com/mike-zeng/pigkit/rpc/log"
	"github.com/mike-zeng/pigkit/rpc/third"
	"net"
)

type ServerTransport interface {
	SetHandlerReq(func(*codec.PigReq) (*codec.PigResponse, error))
	ListenAndServer(ip string,port int)
}


type PigServerTransport struct {
	handlerReq func(context.Context,*codec.PigReq)(*codec.PigResponse, error)
}

func (trans *PigServerTransport) SetHandlerReq(handlerReq func(context.Context,*codec.PigReq) (*codec.PigResponse, error)) {
	trans.handlerReq = handlerReq
}

func (trans *PigServerTransport) ListenAndServer(ip string,port int)  {
	lis, err := net.ListenTCP("tcp",&net.TCPAddr{
		IP:   []byte(ip),
		Port: port,
		Zone: "",
	})
	if err != nil {
		log.DefaultLog.FATAL("network listen error %v",err)
	}
	log.DefaultLog.INFO("server start,listen on port %d",port)
	for {
		conn , err := lis.AcceptTCP()
		if err = conn.SetKeepAlive(true); err != nil {
			log.DefaultLog.FATAL("SetKeepAlive error %v",err)
			return
		}
		if err != nil {
			log.DefaultLog.FATAL("network error %v",err)
		}
		go trans.handleConn(conn)
	}
}

func (trans *PigServerTransport)handleConn(conn net.Conn)  {
	for true {
		frame, err := ReadFrame(conn)
		if err != nil {
			log.DefaultLog.ERROR("read frame error %v",err)
			return
		}
		// encoding frame to req
		request, err := codec.EncodingFrameToRequest(frame)
		if err != nil {
			log.DefaultLog.ERROR("encoding frame error %v",err)
			return
		}
		ctx := codec.WithMetadata(context.Background(), request.MetaData)
		newCtx, span := third.GetTracerClient().StartSpanServer(ctx, request.MethodName)
		// deal req
		resp,err := trans.handlerReq(newCtx,request)
		if span != nil {
			span.Finish()
		}

		if err != nil {
			return
		}
		// decoding resp to frame
		respFrame := codec.DecodingResponseToFrame(resp)
		_, err = conn.Write(respFrame.Bytes())
		if err != nil {
			log.DefaultLog.ERROR("write data to network error %v",err)
		}

	}

}