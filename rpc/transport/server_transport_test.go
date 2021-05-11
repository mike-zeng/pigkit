package transport

import "testing"

func TestPigServerTransport_ListenAndServer(t *testing.T) {
	transport := PigServerTransport{}
	transport.ListenAndServer("127.0.0.1",8181)
}
