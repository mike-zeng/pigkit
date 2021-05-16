package third

import (
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
