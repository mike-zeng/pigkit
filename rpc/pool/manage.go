package pool

import "sync"

type Manage struct {
	poolMap map[string][]NetPool
	lb LoadBalancer
	sync.Mutex
}

func NewManage(lb LoadBalancer) *Manage {
	return &Manage{lb: lb}
}

func (m *Manage) RegisterPool(connInfo ConnInfo)  {
	pool := NewPigNetPool(connInfo)
	pool.Start()
	m.Lock()
	defer m.Unlock()
	if m.poolMap==nil {
		m.poolMap = make(map[string][]NetPool)
	}
	m.poolMap[connInfo.ServerName] = append(m.poolMap[connInfo.ServerName],pool)
}

func (m *Manage) DeletePool(serverName string,connStr string)  {
	//if pools,ok := m.poolMap[serverName];ok {
	//	for _,pool := range pools {
	//		netPool,ok := pool.(PigNetPool)
	//		if ok&&netPool.connInfo.connStr() == connStr{
	//			pool.Close()
	//		}
	//	}
	//}
}

func (m *Manage) GetPool(serverName string) NetPool {
	if pools,ok := m.poolMap[serverName];ok{
		return m.lb.GetPool(pools)
	}
	return nil
}