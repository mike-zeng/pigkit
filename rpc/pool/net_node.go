package pool

import "net"

type netNode struct {
	connStr string // 连接信息ip:port
	Conn net.Conn // 网络连接
	weights int16 // 权重
	lastAccess int64 // 上次访问时间戳： 非临时连接节点该值为0，表示永远不会因为空闲而关闭
	status int32 // 状态
}

const (
	BusyStatus = iota //连接忙碌
	IdleStatus // 连接空闲
	CleanStatus // 连接清理中
	CloseStatus // 连接已关闭
)