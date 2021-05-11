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
	Get() (*NetNode,error)
	// 归还连接
	Return(*NetNode)
	// 刷新连接池
	Update(before func())
	// 启动
	Start()
	// 关闭
	Close()
	// 平滑关闭
	Shutdown()
	// 获取权重
	InitWeights() int
	// 获取当前值
	CurrentWeights() int
	// 设置当前值
	SetCurrentWeights(int)
}

// default impl
type PigNetPool struct {
	connInfo ConnInfo
	connList list.List
	connListBackend []*NetNode
	connNum int32
	weights int
	sync.Mutex
	Status int32
}

func NewPigNetPool(connInfo ConnInfo) *PigNetPool {
	return &PigNetPool{connInfo: connInfo}
}

// 连接池启动
// 根据参数建立连接
func (pool* PigNetPool) Start() {
	for i:=0;i<int(pool.connInfo.MinConnNum);i++{
		node,err := pool.createNode()
		if err != nil{
			// todo 补充日志
		}
		pool.addNode(node)
	}
	if pool.connList.Len() != int(pool.connInfo.MinConnNum) {
		pool.Update(nil)
	}
}

func (pool *PigNetPool) Update(before func()) {
	if before != nil {
		before()
	}

	// 连接重建
	if pool.connNum < pool.connInfo.MinConnNum {
		node, err := pool.createNode()
		if err != nil{
			// todo 补充日志
		}
		pool.addNode(node)
	}

	// 空闲连接清理,清理过程将会lock住
	var next *list.Element
	pool.Lock()
	defer pool.Unlock()
	for e := pool.connList.Front(); e != nil; e = next {
		next = e.Next()
		node := e.Value.(NetNode)
		if (time.Now().Unix()-node.lastAccess) >= pool.connInfo.MaxIdleTime {
			pool.connList.Remove(e)
		}
	}
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

func (pool* PigNetPool) Get()(*NetNode,error) {
	// 加锁
	pool.Lock()
	defer pool.Unlock()

	// 空闲连接不足
	if pool.connList.Len()==0 {
		if pool.connNum >= pool.connInfo.MaxConnNum {
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
		node := front.Value.(*NetNode)
		node.status = BusyStatus
		return node,nil
	}
}

// 归还连接
func (pool* PigNetPool) Return(node *NetNode) {
	pool.Lock()
	defer pool.Unlock()
	pool.connList.PushFront(node)
	node.status = IdleStatus
}

func (pool *PigNetPool) InitWeights()int  {
	return pool.connInfo.InitWeights
}

func (pool *PigNetPool) CurrentWeights() int {
	return pool.weights
}

func (pool *PigNetPool) SetCurrentWeights(num int)  {
	pool.weights += num
}

func (pool* PigNetPool) createNode()(*NetNode,error){
	connStr := pool.connInfo.IpAddr+":"+strconv.Itoa(pool.connInfo.Port)
	conn, err := net.Dial("tcp",connStr)
	if err != nil {
		return nil, err
	}
	node := &NetNode{
		connStr:    connStr,
		Conn:       conn,
	}
	return node,nil
}

func (pool* PigNetPool) addNode(node *NetNode){
	atomic.AddInt32(&pool.connNum,1)
	pool.connList.PushFront(node)
	pool.connListBackend = append(pool.connListBackend,node)
}