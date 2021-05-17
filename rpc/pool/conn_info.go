package pool

// ConnInfo 连接信息
type ConnInfo struct {
	ServiceName string // 服务名
	Addr        string // 连接地址 ip:port
	NodeId      string // 节点id，全局唯一
	InitWeights int // 初始权重
	MaxConnNum  int32 // 最大连接数
	MinConnNum  int32 // 最小连接数
	MaxIdleTime int64 // 最大空闲时间
}