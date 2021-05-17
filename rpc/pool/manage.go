package pool

import (
	"github.com/mike-zeng/pigkit/rpc/third"
	"log"
	"sync"
	"time"
)

// Manage 管理服务所有的连接池
type Manage struct {
	poolMap map[string]map[string]NetPool
	lb LoadBalancer
	sync.Mutex
}

func NewManage(lb LoadBalancer) *Manage {
	return &Manage{lb: lb}
}

// RegisterService 服务注册
func (m *Manage) RegisterService(serviceName string)  {
	// 引入第三方组件ETCD
	etcdService := third.GetEtcdService()
	// 节点初始化触发init函数
	init := func(nodeList []*third.ServiceNode) {
		for _,node := range nodeList {
			m.addNetPool(ConnInfo{
				ServiceName: node.ServiceName,
				Addr:        node.Addr,
				MaxConnNum:  10,
				MinConnNum:  10,
			})
		}
	}
	// 节点上线触发online函数
	online := func(node third.ServiceNode) {
		m.addNetPool(ConnInfo{
			ServiceName: node.ServiceName,
			Addr:        node.Addr,
			InitWeights: 10,
			MaxConnNum:  10,
			MinConnNum:  10,
		})
	}
	// 节点下线触发offline函数
	offline := func(node third.ServiceNode) {
		m.deletePool(node.ServiceName,node.NodeId)
	}
	// 服务发现
	err := etcdService.DetectService(serviceName, init, online, offline)
	if err != nil {
		log.Fatalln(err)
	}
}

// 添加网络连接池
func (m *Manage) addNetPool(connInfo ConnInfo)  {
	pool := NewPigNetPool(connInfo)
	pool.Start()
	go func() {
		time.Sleep(2*time.Duration(connInfo.MaxIdleTime) * time.Millisecond)
		for true {
			pool.Clean()
			time.Sleep(time.Duration(connInfo.MaxIdleTime) * time.Millisecond)
		}
	}()
	m.Lock()
	defer m.Unlock()
	if m.poolMap==nil {
		m.poolMap = make(map[string]map[string]NetPool)
	}
	serviceMap := m.poolMap[connInfo.ServiceName]
	if serviceMap == nil {
		serviceMap = make(map[string]NetPool)
		m.poolMap[connInfo.ServiceName] = serviceMap
	}
	serviceMap[connInfo.NodeId] = pool
}

// deletePool 删除网络连接池
func (m *Manage) deletePool(serverName string,nodeId string)  {
	m.Lock()
	defer m.Unlock()
	pools,ok := m.poolMap[serverName]
	if !ok {
		return
	}
	pool,ok := pools[nodeId]
	if !ok {
		return
	}
	pool.Close()
	delete(pools,nodeId)
}

// GetPool 获取连接池
func (m *Manage) GetPool(serverName string) NetPool {
	if pools,ok := m.poolMap[serverName];ok{
		return m.lb.GetPool(pools)
	}
	return nil
}