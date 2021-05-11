package pool

import "strconv"

// 连接信息
type ConnInfo struct {
	ServerName  string
	IpAddr      string
	Port        int
	InitWeights int
	MaxConnNum  int32
	MinConnNum  int32
	MaxIdleTime int64
}

func (connInfo *ConnInfo) connStr()string  {
	return connInfo.IpAddr+":"+strconv.Itoa(connInfo.Port)
}