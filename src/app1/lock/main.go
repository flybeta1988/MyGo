package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var a string
var once sync.Once

func main() {
	//testBase()
	twoprint()
}

func setup() {
	a = "hello, world"
}

func doprint() {
	once.Do(setup)
	//print(a)
	fmt.Println(a)
}

func twoprint() {
	go doprint()
	go doprint()
	doprint()
}

func testBase() {
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("lock test")
}
