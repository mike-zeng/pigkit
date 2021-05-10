package pool

import (
	"container/list"
	"errors"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type NetPool interface {
	// 获取连接
	Get() (*netNode,error)
	// 归还连接
	Return(*netNode)
	// 刷新连接池
	Update(before func())
	// 启动
	Start()
	// 关闭
	Close()
	// 平滑关闭
	Shutdown()
}

// default impl
type PigNetPool struct {
	connInfo ConnInfo
	connList list.List
	connListBackend []*netNode
	connNum int32
	idleNum int32
	sync.Mutex
	Status int32
}

func NewPigNetPool(connInfo ConnInfo) *PigNetPool {
	return &PigNetPool{connInfo: connInfo}
}

func (pool* PigNetPool) Get()(*netNode,error) {
	// 加锁
	pool.Lock()
	defer pool.Unlock()

	// 空闲连接不足
	if pool.connList.Len()==0 {
		if pool.connNum >= pool.connInfo.maxConnNum {
			return nil,errors.New("没有可用连接")
		}
		node, err := pool.createNode()
		if err != nil {
			return nil,errors.New("没有可用连接")
		}
		node.lastAccess = time.Now().Unix()
		pool.addNode(node)
		return node,nil
	}else{
		front := pool.connList.Front()
		pool.connList.Remove(front)
		if front == nil {
			return nil,errors.New("没有可用连接")
		}
		node := front.Value.(*netNode)
		node.status = BusyStatus
		return node,nil
	}
}

// 归还连接
func (pool* PigNetPool) Return(node *netNode) {
	pool.Lock()
	defer pool.Unlock()
	pool.connList.PushFront(node)
	node.status = IdleStatus
}


// 连接池启动
// 根据参数建立连接
func (pool* PigNetPool) Start() {
	for i:=0;i<int(pool.connInfo.minConnNum);i++{
		node,err := pool.createNode()
		if err != nil{
			// todo 补充日志
		}
		pool.addNode(node)
	}
	if pool.connList.Len() != int(pool.connInfo.minConnNum) {
		pool.Update(nil)
	}
}

func (pool* PigNetPool) Update(before func()) {
	if before != nil {
		before()
	}
	// 连接重建
	if pool.connNum < pool.connInfo.minConnNum {
		node, err := pool.createNode()
		if err != nil{
			// todo 补充日志
		}
		pool.addNode(node)
	}
	// todo 空闲连接清理
}

// 强制关闭所有连接
// 无论当前连接是否空闲，全部关闭
func (pool* PigNetPool) Close() {
	for _,node := range pool.connListBackend {
		node.status = CloseStatus
		_ = node.Conn.Close()
	}
}

func (pool* PigNetPool)Shutdown() {
	for _,node := range pool.connListBackend {
		if node.status == BusyStatus {
			node.status = CleanStatus
		}else{
			node.status = CloseStatus
			_ = node.Conn.Close()
		}
	}
}

func (pool* PigNetPool) createNode()(*netNode,error){
	connStr := pool.connInfo.IpAddr+":"+strconv.Itoa(pool.connInfo.Port)
	conn, err := net.Dial("tcp",connStr)
	if err != nil {
		return nil, err
	}
	node := &netNode{
		connStr:    connStr,
		Conn:       conn,
	}
	return node,nil
}

func (pool* PigNetPool) addNode(node *netNode){
	atomic.AddInt32(&pool.connNum,1)
	pool.connList.PushFront(node)
	pool.connListBackend = append(pool.connListBackend,node)
}