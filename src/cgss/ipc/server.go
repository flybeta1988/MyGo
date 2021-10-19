package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		fmt.Println("ipc.server go Connect -> start... chan len:", len(c))
		for {
			request := <-c
			fmt.Println("ipc.server go Connect request:", request)
			if request == "CLOSE" {
				fmt.Println("ipc.server go Connect -> request close!")
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			fmt.Println("ipc.server go Connect -> req:", req)

			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b)
			fmt.Printf("client.conn len:%d type:%T\n", len(c), c)
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}