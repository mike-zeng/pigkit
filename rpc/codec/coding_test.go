package codec

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestCodingRequestToFrame(t *testing.T) {
	req := PigReq{
		ServerName:        "hello",
		MetaData:          nil,
		SerializationType: 0,
		Content:          "hello world",
	}
	frame := CodingRequestToFrame(&req)
	frameByte := frame.Bytes()
	buffer := bytes.Buffer{}
	buffer.Write(frameByte)

	request, err := EnCodingRequest(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(request)
}
