package main

import (
	rpcserver "./rpc/server"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &rpcserver.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

}
