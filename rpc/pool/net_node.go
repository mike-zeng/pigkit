package pool

import (
	"github.com/mike-zeng/pigkit/rpc/log"
	"net"
	"time"
)

// NetNode 包装网络连接以及相关信息
type NetNode struct {
	connStr string // 连接信息ip:port
	Conn net.Conn // 网络连接
	weights int16 // 权重
	lastAccess int64 // 上次访问时间戳： 非临时连接节点该值为0，表示永远不会因为空闲而关闭
	maxIdle int64
	status int // 状态
}


func (p *NetNode) SetIdleNode()  {
	p.status = IdleStatus
}

func (p *NetNode) SetLastAccess()  {
	p.lastAccess = time.Now().Unix()
}

// CanRecycle 当前节点是否可回收
func (p *NetNode) CanRecycle() bool  {
	return p.IsTemp() &&
		time.Now().Unix() - p.lastAccess >= p.maxIdle
}

func (p *NetNode) IsIdle()bool  {
	return p.status == IdleStatus
}

func (p *NetNode) IsTemp()bool  {
	return p.status == TempStatus
}

func (p *NetNode) IsBusy()bool  {
	return p.status == BusyStatus
}

func (p *NetNode) IsClean()bool  {
	return p.status == CleanStatus
}

func (p *NetNode) IsClose()bool  {
	return p.status == CloseStatus
}

func (p *NetNode) Close() {
	p.status = CleanStatus
	err := p.Conn.Close()
	if err !=nil {
		log.DefaultLog.WARNING("close net conn error: %s",err)
		return
	}
	p.status = CloseStatus
}

const (
	IdleStatus  = iota// 连接空闲
	TempStatus // 临时状态
	BusyStatus  //连接忙碌
	CleanStatus // 连接清理中
	CloseStatus // 连接已关闭
)