package client

import "time"

type Options struct {
	serviceName string // service name
	method string // method name
	timeout time.Duration  // timeout
	SerializationType string
}