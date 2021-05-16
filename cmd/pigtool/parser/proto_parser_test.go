package parser

import (
	"testing"
)

func TestProtoParser_ReadFromIdl(t *testing.T) {
	parser := NewProtoParser("../hello.proto")
	println(parser.GetServiceName())
}
