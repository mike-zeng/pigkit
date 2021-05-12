package test

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":8181")
	if err != nil {
		t.Error(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		t.Error(err)
	}
	buf := make([]byte,5)
	//content, err := conn.Read(buf)
	//content, err := io.ReadAtLeast(conn,buf,4)
	content, err := io.ReadFull(conn, buf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(buf)
	fmt.Println(content)
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil {
		t.Error(err)
	}
	_, err = conn.Write([]byte("he44444444444444444444444444"))
	if err != nil {
		t.Error(err)
	}
	time.Sleep(100*time.Second)
}

func TestJson(t *testing.T) {

}