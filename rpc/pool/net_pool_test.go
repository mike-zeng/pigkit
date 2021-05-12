package pool

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func startDemoServer() {
	func() {
		listener, err := net.Listen("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}
		for {
			_, err := listener.Accept()
			if err != nil {
				log.Print(err) // 例如，连接终止
				continue
			}
			fmt.Println("连接成功")
		}
	}()
}

func TestPigNetPool_Get(t *testing.T) {
	//go startDemoServer()
	//pool := NewPigNetPool(ConnInfo{
	//	ServiceName:  "test",
	//	IpAddr:      "localhost",
	//	Port:        8000,
	//	MaxConnNum:  10,
	//	MinConnNum:  10,
	//	MaxIdleTime: 2000,
	//})
	//pool.Start()
	//var count int32 = 0
	//
	//wait := sync.WaitGroup{}
	//wait.Add(1)
	//for i:=0;i<10;i++ {
	//	go func() {
	//		wait.Wait()
	//		for{
	//			// 获取连接
	//			conn, err := pool.Get()
	//			if err != nil {
	//				fmt.Println(err)
	//			}
	//			// 归还连接
	//			pool.Return(conn)
	//			atomic.AddInt32(&count,1)
	//		}
	//	}()
	//}
	//wait.Done()
	//fmt.Println("go:")
	//time.Sleep(1*time.Second)
	//fmt.Printf("count=%d\n",count)
}
