package client

import "time"

type Options struct {
	ServiceName       string        // service name
	Method            string        // Method name
	Timeout           time.Duration // Timeout
	SerializationType string
}