package third

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_etcdService_RegService(t *testing.T) {
	ser := GetEtcdService()
	err := ser.RegService("12222222", "test", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(100*time.Second)
}

func Test_etcdService_GetService(t *testing.T) {
	ser := GetEtcdService()
	init := func(nodeList []*ServiceNode) {
		fmt.Println("init")
		for _, node := range nodeList {
			fmt.Println(*node)
		}
	}
	online := func(node ServiceNode) {
		fmt.Println("online")
		fmt.Println(node)
	}
	offline := func(node ServiceNode) {
		fmt.Println("offline")
		fmt.Println(node)
	}
	err := ser.DetectService("test",init ,online ,offline)
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(100*time.Second)
}