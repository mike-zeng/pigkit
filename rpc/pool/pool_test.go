package pool

import (
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	// 1. 创建服务器
	ser, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	for true {
		accept, err := ser.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("已连接")
		go func() {
			for {
				var bytes []byte
				_, err2 := accept.Read(bytes)
				if err2 != nil {
					log.Fatalln(err2)
				}

			}
		}()
	}
	log.Println("service start...")
	time.Sleep(100*time.Second)
}

// benchmark for pool
func BenchmarkPool(b *testing.B) {
	// 2. 创建连接池
	manage := NewManage(DefaultLoadBalancer())
	manage.addNetPool(ConnInfo{
		ServiceName: "test_service",
		Addr:        "127.0.0.1:8081",
		NodeId:      "5201314",
		MaxConnNum:  10,
		MinConnNum:  20,
	})
	b.ResetTimer()
	// 3. 并发读取
	group := sync.WaitGroup{}
	group.Add(b.N)
	for k := 0; k < b.N; k++ {
		go func() {
			conn, err := manage.GetPool("test_service").Get()
			if err != nil {
				log.Fatalln(err)
			}
			manage.GetPool("test_service").Return(conn)
			group.Done()
		}()
	}
	group.Wait()
}