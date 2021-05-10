package pool

import (
	"testing"
)

func TestManage_GetPool(t *testing.T) {
	go startDemoServer()
	info := ConnInfo{
		ServerName:  "test",
		IpAddr:      "localhost",
		Port:        8000,
		maxConnNum:  10,
		minConnNum:  10,
		maxIdleTime: 2000,
	}
	manage := NewManage(LoadBalancer{})
	manage.RegisterPool(info)
	pool := manage.GetPool("test")
	conn, err := pool.Get()
	if err != nil {
		t.Error(err)
	}
	pool.Return(conn)
}
