package server

type ServiceDesc struct {
	ServiceName string
	Methods     map[string]Handler
	HandlerType interface{}
}