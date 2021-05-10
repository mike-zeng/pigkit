package pool
// 连接信息
type ConnInfo struct {
	ServerName string
	IpAddr string
	Port int
	maxConnNum int32
	minConnNum int32
	maxIdleTime int32
}
