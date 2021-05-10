package pool

import "strconv"

// 连接信息
type ConnInfo struct {
	ServerName string
	IpAddr string
	Port int
	initWeights int
	maxConnNum int32
	minConnNum int32
	maxIdleTime int64
}

func (connInfo *ConnInfo) connStr()string  {
	return connInfo.IpAddr+":"+strconv.Itoa(connInfo.Port)
}