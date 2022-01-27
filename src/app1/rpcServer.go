package main

import (
	rpcserver "./rpc/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(rpcserver.Arith)
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal("rpc register error:", err)
	}

	rpc.HandleHTTP()
	l, e := net.Listen("tcp",  ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
