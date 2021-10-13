package ipc

import "testing"

type EchoServer struct {
	//Server
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200", "ok"}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	/*resp1, err := client1.Call("get", "aaa")
	resp2, err := client2.Call("get", "bbb")*/

	client1.Close()
	client2.Close()
}
