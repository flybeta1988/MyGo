package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i ++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

//输出的结果无序且个数会少于10
func main() {
	go say("b")
	say("a")
}
