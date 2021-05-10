package pool

import "sync"

type Manage struct {
	poolMap map[string][]NetPool
	sync.Mutex
}

func (m *Manage) RegisterPool(connInfo ConnInfo)  {
	m.Lock()
	defer m.Unlock()
	pool := NewPigNetPool(connInfo)
	pool.Start()
	m.poolMap[connInfo.ServerName] = append(m.poolMap[connInfo.ServerName],pool)
}

func (m *Manage) DeletePool(serverName string,pool NetPool)  {

}

func (m *Manage) GetPool()  {

}