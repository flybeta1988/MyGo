package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error)  {
	req := &Request{method, params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	fmt.Println("before client.conn.len:", len(client.conn))

	jsonStr := string(b)
	select {
	case client.conn <- jsonStr:
		fmt.Println("client.conn write ok")
		fmt.Printf("client.conn len:%d type:%T\n", len(client.conn), client.conn)
	default:
		fmt.Println("client.conn write failed!")
	}

	str := <-client.conn //等待返回值
	fmt.Println("after client.conn.len:", len(client.conn), " str:", str)

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
