package main

import (
	"fmt"
	"sync"
	"time"
)

var a string
var once sync.Once

func main() {
	twoPrint()
	time.Sleep(2 * time.Second)
}

func setup() {
	a = "hello once"
	fmt.Println("init a")
}

func doPrint() {
	once.Do(setup)
	//setup()
	fmt.Println(a)
}

func twoPrint() {
	go doPrint()
	go doPrint()
}
