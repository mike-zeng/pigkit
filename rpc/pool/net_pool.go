package pool

import (
	"github.com/mike-zeng/pigkit/rpc/log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type NetPool interface {
	// Get 获取连接
	Get() (*NetNode,error)
	// Return 归还连接
	Return(*NetNode)
	// Info 连接池信息
	Info()ConnInfo
	// Start 启动
	Start()
	// Close 关闭
	Close()
	// Shutdown 平滑关闭
	Shutdown()
	// InitWeights 获取权重
	InitWeights() int
	// CurrentWeights 获取当前值
	CurrentWeights() int
	// SetCurrentWeights 设置当前值
	SetCurrentWeights(int)
	// Clean 空闲连接回收
	Clean()
}

// PigNetPool default impl
type PigNetPool struct {
	connInfo ConnInfo
	connNum int32
	weights int
	sync.Mutex
	Status int32
	queue NodeQueue
}

func NewPigNetPool(connInfo ConnInfo) *PigNetPool {
	return &PigNetPool{connInfo: connInfo,queue: NewBlockNodeQueue(int(connInfo.MinConnNum))}
}

func (pool* PigNetPool) Info() ConnInfo{
	return pool.Info()
}
// Start 连接池启动
// 根据参数建立连接
func (pool* PigNetPool) Start() {
	for i:=0;i<int(pool.connInfo.MinConnNum);i++{
		node,err := pool.createNode()
		if err != nil{
			go func() {
				for i := 0; i < 10; i++ {
					node,err := pool.createNode()
					if err != nil {
						time.Sleep(1*time.Second)
						continue
					}
					pool.addNode(node)
					return
				}
				log.DefaultLog.WARNING("init conn error:",err)
			}()
			return
		}
		pool.addNode(node)
	}
}


// Close 强制关闭所有连接
// 无论当前连接是否空闲，全部关闭
func (pool* PigNetPool) Close() {

}

func (pool* PigNetPool)Shutdown() {
	panic("no use")
}

func (pool* PigNetPool) Get()(*NetNode,error) {
	node := pool.queue.Get()
	if node == nil {
		for pool.connInfo.MaxConnNum > pool.connNum {
			if atomic.CompareAndSwapInt32(&pool.connNum, pool.connNum, pool.connNum+1) {
				// 创建空闲连接并返回
				createNode, err := pool.createNode()
				if err != nil {
					atomic.AddInt32(&pool.connNum,-1)
					return pool.queue.Take(),nil
				}
				createNode.status = IdleStatus
				return createNode,nil
			}
		}
		return pool.queue.Take(),nil
	}
	return node,nil
}

// Return 归还连接
func (pool* PigNetPool) Return(node *NetNode) {
	if node.IsTemp() {
		node.SetLastAccess()
	}
	pool.queue.Put(node)
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
	conn, err := net.Dial("tcp",pool.connInfo.Addr)
	if err != nil {
		return nil, err
	}
	node := &NetNode{
		connStr:    pool.connInfo.Addr,
		Conn:       conn,
		maxIdle: pool.connInfo.MaxIdleTime,
	}
	return node,nil
}

func (pool* PigNetPool) addNode(node *NetNode){
	pool.queue.Put(node)
}

func (pool* PigNetPool) Clean()  {
	pool.queue.Clean()
}