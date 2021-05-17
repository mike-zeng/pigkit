package third

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/mike-zeng/pigkit/rpc/config"
	"go.etcd.io/etcd/clientv3"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

func init() {

}

type etcdService struct {
	etcdClient *clientv3.Client
	lock sync.Mutex
	serverList    map[string]ServiceNode
}

type ServiceNode struct {
	ServiceName string
	Addr        string
	NodeId      string
}

func newServiceNode(key string,val string) *ServiceNode {
	serviceName := getServiceName(key)
	nodeId := getNodeId(key)
	addr := getAddr(val)
	return &ServiceNode{ServiceName: serviceName, NodeId: nodeId, Addr: addr}
}

const Prefix = "/pig_services/"

// 单例模式
var onceForEtcd sync.Once
var service *etcdService
func GetEtcdService() *etcdService {
	onceForEtcd.Do(func() {
		client,err := clientv3.New(clientv3.Config{
			Endpoints:   config.GetConfig().PigServer.Etcd.Hosts,
			DialTimeout: time.Duration(config.GetConfig().PigServer.TimeOut)*time.Millisecond,
		})
		if err != nil {
			log.Fatalln(err)
		}
		service = &etcdService{
			etcdClient: client,
		}
	})
	return service
}

// RegService 服务注册
func (ser *etcdService)RegService(nodeId,serviceName string, address string) error {
	client := ser.etcdClient
	kv := clientv3.NewKV(client)
	ctx := context.Background()
	lease := clientv3.NewLease(client)

	//设置租约过期时间为5秒
	leaseRes, err := clientv3.NewLease(client).Grant(ctx, 5)
	if err != nil {
		return err
	}
	_, err = kv.Put(context.Background(), Prefix + serviceName + "/" + nodeId, address, clientv3.WithLease(leaseRes.ID)) //把服务的key绑定到租约下面
	if err != nil {
		return err
	}
	keepaliveRes, err := lease.KeepAlive(context.TODO(), leaseRes.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case ret := <-keepaliveRes:
				if ret != nil {

				}
			}
		}
	}()
	return err
}

// DetectService 服务发现
func (ser *etcdService)DetectService(serviceName string,init func(nodeList []*ServiceNode),
	online func(node ServiceNode),offline func(node ServiceNode)) error {
	client := ser.etcdClient
	serviceKey := Prefix + serviceName
	resp, err := client.Get(context.Background(), serviceKey, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	nodeList := ser.extractServiceNodeList(resp)
	init(nodeList)
	go ser.watcher(serviceKey,offline,online)
	return nil
}


// 将获取到的prefix目录下所有内容存入list并返回
func (ser *etcdService)extractServiceNodeList(resp *clientv3.GetResponse) []*ServiceNode {
	var nodeList []*ServiceNode
	if resp == nil || resp.Kvs == nil {
		return nodeList
	}
	for _,keyValue := range resp.Kvs {
		value := keyValue.Value
		key := keyValue.Key
		serviceName := getServiceName(string(key))
		nodeId := getNodeId(string(key))
		addr := getAddr(string(value))
		nodeList = append(nodeList, &ServiceNode{
			ServiceName: serviceName,
			Addr:        addr,
			NodeId:      nodeId,
		})
	}
	return nodeList
}

func getAddr(value string)string {
	if value == "" {
		return ""
	}
	split := strings.Split(value,"#")
	return split[0]
}

// watch负责将监听到的put、delete请求存放到指定list
func (ser *etcdService)watcher(prefix string,offline func(nodeList ServiceNode),online func(nodeList ServiceNode)) {
	client := ser.etcdClient
	rch := client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				online(*newServiceNode(string(ev.Kv.Key),string(ev.Kv.Value)))
			case mvccpb.DELETE:
				offline(*newServiceNode(string(ev.Kv.Key),string(ev.Kv.Value)))
			}
		}
	}
}

func getWeight(value string)int {
	split := strings.Split(value,"#")
	if len(split) == 1 {
		return 0
	}
	num, err := strconv.Atoi(split[1])
	if err != nil {
		return 0
	}
	return num
}


func getNodeId(key string)string {
	split := strings.Split(key,"/")
	fmt.Println(split)
	if len(split) != 4 {
		return ""
	}
	return split[3]
}

func getServiceName(key string)string {
	split := strings.Split(key,"/")
	if len(split) != 4 {
		return ""
	}
	return split[2]
}
