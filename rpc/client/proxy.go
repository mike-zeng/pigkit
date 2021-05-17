package client

type Proxy interface {
	GetProxyServiceName()string
	RegisterCallback(client Client)
}
