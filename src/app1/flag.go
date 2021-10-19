package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//方法1
	var foo string
	flag.StringVar(&foo, "foo", "xxx", "any other things")

	//方法2
	act := flag.String("act", "add", "this is act type")
	flag.Parse()

	//方法3
	args := os.Args
	arg1 := args[1]
	arg1arr := strings.Split(arg1, "=")

	fmt.Println(*act)
	fmt.Println(foo)
	fmt.Println(args)
	fmt.Println(arg1arr[1])
}
